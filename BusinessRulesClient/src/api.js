const BASE_URL = 'http://localhost:3000';

export const apiCall = async (endpoint, payload = {}) => {
    try {
        const response = await fetch(`${BASE_URL}${endpoint}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(payload)
        });

        const data = await response.json();
        return data;
    } catch (error) {
        return { message: error.message || "Network error. Please try again later.", error: true };
    }
};

export const CategoryAPI = {
    getAll: () => apiCall('/v1/category/get-all', {}),
    create: (name) => apiCall('/v1/category/create', { name })
};

export const AttributeAPI = {
    getAll: () => apiCall('/v1/attributes/get-all', {}),
    create: (name, dataType) => apiCall('/v1/attribute/create', { name, dataType })
};

export const AssignmentAPI = {
    getCategoryWiseCommonAttributes: (categoryIds) => apiCall('/v1/assignment/get-category-wise-common-attributes', { categoryIds }),
    change: (assignCategories, assignAttributes, unassignCategories, unassignAttributes) => apiCall('/v1/assignment/change', {
        assign: { categoryIds: assignCategories, attributeIds: assignAttributes },
        unassign: { categoryIds: unassignCategories, attributeIds: unassignAttributes }
    })
};

export const ProductAPI = {
    getAll: () => apiCall('/v1/product/get-all', {}),
    getById: (productId) => apiCall('/v1/product/get-by-id', { productId }),
    // When creating a new product, use "" string for productId
    upsert: (categoryId, productId, data) => apiCall('/v1/product/upsert', { categoryId, productId, data })
};

export const FormulaAPI = {
    getAll: () => apiCall('/v1/formula/get-all', {}),
    create: (categoryId, targetAttribute, formula) => apiCall('/v1/formula/create', { categoryId, targetAttribute, formula })
};
