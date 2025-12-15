// Automated API Testing Script for Order Management System
// Run with: node test_endpoints.js

const BASE_URL = 'http://localhost:8080/api/v1';
let authToken = '';
let userId = '';
let productId = '';

// Color codes for console output
const colors = {
    reset: '\x1b[0m',
    green: '\x1b[32m',
    red: '\x1b[31m',
    yellow: '\x1b[33m',
    blue: '\x1b[34m',
    cyan: '\x1b[36m'
};

// Helper function to make requests
async function makeRequest(endpoint, method = 'GET', body = null, useAuth = false) {
    const headers = {
        'Content-Type': 'application/json'
    };

    if (useAuth && authToken) {
        headers['Authorization'] = `Bearer ${authToken}`;
    }

    const options = {
        method,
        headers
    };

    if (body) {
        options.body = JSON.stringify(body);
    }

    try {
        const response = await fetch(`${BASE_URL}${endpoint}`, options);
        const data = await response.json();
        return { status: response.status, data, ok: response.ok };
    } catch (error) {
        return { status: 0, error: error.message, ok: false };
    }
}

// Test helper function
function logTest(testName, passed, message = '') {
    const icon = passed ? '✓' : '✗';
    const color = passed ? colors.green : colors.red;
    console.log(`${color}${icon} ${testName}${colors.reset}${message ? ' - ' + message : ''}`);
}

// Generate random email for testing
function generateTestEmail() {
    const timestamp = Date.now();
    return `test${timestamp}@example.com`;
}

// Main test suite
async function runTests() {
    console.log(`\n${colors.cyan}═══════════════════════════════════════════════════════${colors.reset}`);
    console.log(`${colors.cyan}    Order Management System - API Tests${colors.reset}`);
    console.log(`${colors.cyan}═══════════════════════════════════════════════════════${colors.reset}\n`);

    const testEmail = generateTestEmail();
    const testPassword = 'Test@123456';

    // Test 1: Health Check
    console.log(`${colors.blue}[1] Testing Health Check Endpoint...${colors.reset}`);
    const healthCheck = await makeRequest('/hello');
    logTest('GET /api/v1/hello', healthCheck.ok, healthCheck.data?.message);

    // Test 2: User Registration
    console.log(`\n${colors.blue}[2] Testing User Registration...${colors.reset}`);
    const registerData = {
        name: 'Test User',
        email: testEmail,
        password: testPassword
    };
    const register = await makeRequest('/auth/register', 'POST', registerData);
    logTest('POST /api/v1/auth/register', register.ok, register.data?.message);
    if (register.ok && register.data?.data?.id) {
        userId = register.data.data.id;
        console.log(`   User ID: ${userId}`);
    }

    // Test 3: User Login
    console.log(`\n${colors.blue}[3] Testing User Login...${colors.reset}`);
    const loginData = {
        email: testEmail,
        password: testPassword
    };
    const login = await makeRequest('/auth/login', 'POST', loginData);
    logTest('POST /api/v1/auth/login', login.ok, login.data?.message);
    if (login.ok && login.data?.token) {
        authToken = login.data.token;
        console.log(`   Token received: ${authToken.substring(0, 50)}...`);
    }

    if (!authToken) {
        console.log(`\n${colors.red}❌ Tests stopped: Authentication failed${colors.reset}\n`);
        return;
    }

    // Test 4: Get All Users
    console.log(`\n${colors.blue}[4] Testing Get All Users...${colors.reset}`);
    const users = await makeRequest('/users', 'GET', null, true);
    logTest('GET /api/v1/users', users.ok, `Found ${users.data?.count || 0} users`);

    // Test 5: Get User by ID
    console.log(`\n${colors.blue}[5] Testing Get User by ID...${colors.reset}`);
    const user = await makeRequest(`/users/${userId}`, 'GET', null, true);
    logTest(`GET /api/v1/users/${userId}`, user.ok, user.data?.data?.email);

    // Test 6: Update User
    console.log(`\n${colors.blue}[6] Testing Update User...${colors.reset}`);
    const updateUserData = {
        name: 'Test User Updated'
    };
    const updateUser = await makeRequest(`/users/${userId}`, 'PUT', updateUserData, true);
    logTest(`PUT /api/v1/users/${userId}`, updateUser.ok, updateUser.data?.message);

    // Test 7: Create Product
    console.log(`\n${colors.blue}[7] Testing Create Product...${colors.reset}`);
    const productData = {
        name: 'Test Laptop',
        description: 'High performance test laptop',
        price: 1299.99,
        quantity: 50,
        category: 'Electronics',
        sku: `TEST-LAP-${Date.now()}`
    };
    const createProduct = await makeRequest('/products', 'POST', productData, true);
    logTest('POST /api/v1/products', createProduct.ok, createProduct.data?.message);
    if (createProduct.ok && createProduct.data?.data?.id) {
        productId = createProduct.data.data.id;
        console.log(`   Product ID: ${productId}`);
    }

    // Test 8: Get All Products
    console.log(`\n${colors.blue}[8] Testing Get All Products...${colors.reset}`);
    const products = await makeRequest('/products', 'GET', null, true);
    logTest('GET /api/v1/products', products.ok, `Found ${products.data?.count || 0} products`);

    // Test 9: Get Product by ID
    if (productId) {
        console.log(`\n${colors.blue}[9] Testing Get Product by ID...${colors.reset}`);
        const product = await makeRequest(`/products/${productId}`, 'GET', null, true);
        logTest(`GET /api/v1/products/${productId}`, product.ok, product.data?.data?.name);
    }

    // Test 10: Search Products
    console.log(`\n${colors.blue}[10] Testing Search Products...${colors.reset}`);
    const search = await makeRequest('/products/search?q=laptop', 'GET', null, true);
    logTest('GET /api/v1/products/search?q=laptop', search.ok, `Found ${search.data?.count || 0} results`);

    // Test 11: Get Products by Category
    console.log(`\n${colors.blue}[11] Testing Get Products by Category...${colors.reset}`);
    const category = await makeRequest('/products/category/Electronics', 'GET', null, true);
    logTest('GET /api/v1/products/category/Electronics', category.ok, `Found ${category.data?.count || 0} products`);

    // Test 12: Update Product (Full)
    if (productId) {
        console.log(`\n${colors.blue}[12] Testing Update Product (PUT)...${colors.reset}`);
        const updateProductData = {
            name: 'Test Gaming Laptop',
            description: 'Updated gaming laptop',
            price: 1599.99,
            quantity: 30
        };
        const updateProduct = await makeRequest(`/products/${productId}`, 'PUT', updateProductData, true);
        logTest(`PUT /api/v1/products/${productId}`, updateProduct.ok, updateProduct.data?.message);
    }

    // Test 13: Update Product (Partial)
    if (productId) {
        console.log(`\n${colors.blue}[13] Testing Update Product (PATCH)...${colors.reset}`);
        const patchProductData = {
            price: 1499.99
        };
        const patchProduct = await makeRequest(`/products/${productId}`, 'PATCH', patchProductData, true);
        logTest(`PATCH /api/v1/products/${productId}`, patchProduct.ok, patchProduct.data?.message);
    }

    // Test 14: Update Product Quantity
    if (productId) {
        console.log(`\n${colors.blue}[14] Testing Update Product Quantity...${colors.reset}`);
        const quantityData = {
            quantity: 100
        };
        const updateQuantity = await makeRequest(`/products/${productId}/quantity`, 'PUT', quantityData, true);
        logTest(`PUT /api/v1/products/${productId}/quantity`, updateQuantity.ok, updateQuantity.data?.message);
    }

    // Test 15: Delete Product
    if (productId) {
        console.log(`\n${colors.blue}[15] Testing Delete Product...${colors.reset}`);
        const deleteProduct = await makeRequest(`/products/${productId}`, 'DELETE', null, true);
        logTest(`DELETE /api/v1/products/${productId}`, deleteProduct.ok, deleteProduct.data?.message);
    }

    // Test 16: Delete User (cleanup)
    console.log(`\n${colors.blue}[16] Testing Delete User (Cleanup)...${colors.reset}`);
    const deleteUser = await makeRequest(`/users/${userId}`, 'DELETE', null, true);
    logTest(`DELETE /api/v1/users/${userId}`, deleteUser.ok, deleteUser.data?.message);

    // Summary
    console.log(`\n${colors.cyan}═══════════════════════════════════════════════════════${colors.reset}`);
    console.log(`${colors.green}✓ All tests completed!${colors.reset}`);
    console.log(`${colors.cyan}═══════════════════════════════════════════════════════${colors.reset}\n`);
}

// Run the tests
console.log(`${colors.yellow}Starting API tests...${colors.reset}`);
console.log(`${colors.yellow}Make sure the server is running on ${BASE_URL}${colors.reset}\n`);

runTests().catch(error => {
    console.error(`\n${colors.red}❌ Test suite failed:${colors.reset}`, error);
    process.exit(1);
});
