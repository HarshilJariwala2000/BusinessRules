import React, { useState, useEffect } from 'react';
import { CategoryAPI, AssignmentAPI } from '../api';
import { AlertCircle, Link as LinkIcon, Loader2 } from 'lucide-react';

export default function AssignmentPage() {
    const [categories, setCategories] = useState([]);
    const [selectedCategories, setSelectedCategories] = useState([]);

    const [attributes, setAttributes] = useState([]);
    const [loadingCategories, setLoadingCategories] = useState(false);
    const [loadingAttributes, setLoadingAttributes] = useState(false);
    const [error, setError] = useState('');

    useEffect(() => {
        fetchCategories();
    }, []);

    useEffect(() => {
        if (selectedCategories.length > 0) {
            fetchAssignments();
        } else {
            setAttributes([]);
        }
    }, [selectedCategories]);

    const fetchCategories = async () => {
        setLoadingCategories(true);
        const res = await CategoryAPI.getAll();
        if (res.message === 'success') {
            setCategories(res.data || []);
        } else {
            setError(res.message);
        }
        setLoadingCategories(false);
    };

    const fetchAssignments = async () => {
        setLoadingAttributes(true);
        const res = await AssignmentAPI.getCategoryWiseCommonAttributes(selectedCategories);
        if (res.message === 'success') {
            setAttributes(res.data || []);
            setError('');
        } else {
            setError(res.message);
        }
        setLoadingAttributes(false);
    };

    const toggleCategory = (categoryId) => {
        setSelectedCategories(prev =>
            prev.includes(categoryId)
                ? prev.filter(id => id !== categoryId)
                : [...prev, categoryId]
        );
    };

    const toggleAttributeAssignment = async (attributeId, isCurrentlyAssigned) => {
        const isAssigning = !isCurrentlyAssigned;

        // Optimistic UI update
        setAttributes(prev => prev.map(attr =>
            attr.id === attributeId ? { ...attr, assigned: isAssigning } : attr
        ));

        const assignCats = isAssigning ? selectedCategories : [];
        const assignAttrs = isAssigning ? [attributeId] : [];
        const unassignCats = !isAssigning ? selectedCategories : [];
        const unassignAttrs = !isAssigning ? [attributeId] : [];

        const res = await AssignmentAPI.change(assignCats, assignAttrs, unassignCats, unassignAttrs);
        if (res.message !== 'success') {
            setError(res.message);
            // Revert if failed
            setAttributes(prev => prev.map(attr =>
                attr.id === attributeId ? { ...attr, assigned: isCurrentlyAssigned } : attr
            ));
        } else {
            // Re-fetch to ensure sync with server
            fetchAssignments();
        }
    };

    return (
        <div className="page-container">
            <div className="flex justify-between items-center" style={{ marginBottom: '2rem' }}>
                <h1>Category-Attribute Assignments</h1>
            </div>

            {error && (
                <div className="alert alert-error">
                    <AlertCircle size={20} />
                    <span>{error}</span>
                </div>
            )}

            <div className="flex gap-4">
                {/* Left Column: Categories */}
                <div className="card w-full" style={{ maxWidth: '350px' }}>
                    <div className="flex items-center gap-2 mb-4">
                        <h3>Categories</h3>
                        {loadingCategories && <Loader2 size={16} className="animate-spin" />}
                    </div>

                    <div className="flex flex-col gap-2">
                        {categories.length === 0 && !loadingCategories && (
                            <p className="badge badge-neutral">No categories available.</p>
                        )}
                        {categories.map(cat => (
                            <label key={cat.id} className="checkbox-container" style={{ padding: '0.5rem', background: selectedCategories.includes(cat.id) ? 'rgba(99, 102, 241, 0.1)' : 'transparent', borderRadius: '0.5rem' }}>
                                <input
                                    type="checkbox"
                                    checked={selectedCategories.includes(cat.id)}
                                    onChange={() => toggleCategory(cat.id)}
                                />
                                <span style={{ fontWeight: selectedCategories.includes(cat.id) ? '600' : '400' }}>
                                    {cat.name}
                                </span>
                            </label>
                        ))}
                    </div>
                </div>

                {/* Right Column: Attributes */}
                <div className="card flex-1">
                    <div className="flex items-center gap-2 mb-4">
                        <h3>Attributes for Selected Categories</h3>
                        {loadingAttributes && <Loader2 size={16} className="animate-spin" />}
                    </div>

                    {selectedCategories.length === 0 ? (
                        <div style={{ textAlign: 'center', padding: '3rem', color: '#8b949e' }}>
                            <LinkIcon size={48} style={{ margin: '0 auto 1rem', opacity: 0.5 }} />
                            <p>Select one or more categories on the left to view and manage assignments.</p>
                        </div>
                    ) : (
                        <div>
                            {attributes.length === 0 && !loadingAttributes ? (
                                <p className="badge badge-neutral">No attributes found for these categories.</p>
                            ) : (
                                <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(250px, 1fr))', gap: '1rem' }}>
                                    {attributes.map(attr => (
                                        <div key={attr.id} className="flex items-center justify-between" style={{ padding: '1rem', background: '#21262d', borderRadius: '0.5rem', border: '1px solid #30363d' }}>
                                            <div className="flex flex-col">
                                                <span style={{ fontWeight: '500' }}>{attr.name}</span>
                                                <span style={{ fontSize: '0.75rem', color: '#8b949e' }}>Type: {attr.dataType}</span>
                                            </div>
                                            <label className="checkbox-container">
                                                <input
                                                    type="checkbox"
                                                    checked={attr.assigned}
                                                    onChange={() => toggleAttributeAssignment(attr.id, attr.assigned)}
                                                />
                                            </label>
                                        </div>
                                    ))}
                                </div>
                            )}
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
}
