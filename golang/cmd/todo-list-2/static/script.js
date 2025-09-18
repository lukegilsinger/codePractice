// script.js (COMPLETELY UPDATED with Authentication)
const API_BASE = '/api';
let categories = [];
let todos = [];
let selectedCategoryFilter = null;
let currentUser = null;

// ===================================================================
// AUTHENTICATION & INITIALIZATION
// ===================================================================

document.addEventListener('DOMContentLoaded', async () => {
    checkAuthState();
});

function checkAuthState() {
    const token = localStorage.getItem('todo_token');
    if (token) {
        // Validate token and load user data
        loadUserData();
    } else {
        showAuthSection();
    }
}

async function loadUserData() {
    try {
        const response = await apiCall('/auth/me', 'GET');
        if (response.ok) {
            currentUser = await response.json();
            showAppSection();
            await loadCategories();
            await loadTodos();
        } else {
            // Token invalid, show login
            localStorage.removeItem('todo_token');
            showAuthSection();
        }
    } catch (error) {
        console.error('Error loading user data:', error);
        localStorage.removeItem('todo_token');
        showAuthSection();
    }
}

function showAuthSection() {
    document.getElementById('auth-section').style.display = 'flex';
    document.getElementById('app-section').classList.remove('show');
}

function showAppSection() {
    document.getElementById('auth-section').style.display = 'none';
    document.getElementById('app-section').classList.add('show');
    updateUserDisplay();
}

function updateUserDisplay() {
    if (currentUser) {
        document.getElementById('user-name').textContent = currentUser.username;
        document.getElementById('user-email').textContent = currentUser.email;
        document.getElementById('user-avatar').textContent = currentUser.username.charAt(0).toUpperCase();
    }
}

// ===================================================================
// API HELPER FUNCTIONS
// ===================================================================

async function apiCall(endpoint, method = 'GET', data = null) {
    const token = localStorage.getItem('todo_token');
    const headers = {
        'Content-Type': 'application/json',
    };
    
    if (token) {
        headers.Authorization = `Bearer ${token}`;
    }
    
    const config = {
        method,
        headers,
    };
    
    if (data) {
        config.body = JSON.stringify(data);
    }
    
    return fetch(`${API_BASE}${endpoint}`, config);
}

// ===================================================================
// AUTHENTICATION HANDLERS
// ===================================================================

function switchAuthTab(tab) {
    // Update tabs
    document.querySelectorAll('.auth-tab').forEach(t => t.classList.remove('active'));
    document.querySelector(`[onclick="switchAuthTab('${tab}')"]`).classList.add('active');
    
    // Update forms
    document.querySelectorAll('.auth-form').forEach(f => f.classList.remove('active'));
    document.getElementById(`${tab}-form`).classList.add('active');
    
    // Clear errors
    clearAuthErrors();
}

function clearAuthErrors() {
    document.getElementById('login-error').textContent = '';
    document.getElementById('register-error').textContent = '';
    document.getElementById('register-success').textContent = '';
}

async function handleLogin(event) {
    event.preventDefault();
    clearAuthErrors();
    
    const username = document.getElementById('login-username').value.trim();
    const password = document.getElementById('login-password').value;
    
    try {
        const response = await fetch(`${API_BASE}/auth/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, password }),
        });
        
        if (response.ok) {
            const data = await response.json();
            localStorage.setItem('todo_token', data.token);
            currentUser = data.user;
            showAppSection();
            await loadCategories();
            await loadTodos();
            
            // Clear form
            document.getElementById('login-form').reset();
        } else {
            const error = await response.text();
            document.getElementById('login-error').textContent = error || 'Login failed';
        }
    } catch (error) {
        console.error('Login error:', error);
        document.getElementById('login-error').textContent = 'Network error. Please try again.';
    }
}

async function handleRegister(event) {
    event.preventDefault();
    clearAuthErrors();
    
    const username = document.getElementById('register-username').value.trim();
    const email = document.getElementById('register-email').value.trim();
    const password = document.getElementById('register-password').value;
    
    try {
        const response = await fetch(`${API_BASE}/auth/register`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, email, password }),
        });
        
        if (response.ok) {
            const data = await response.json();
            localStorage.setItem('todo_token', data.token);
            currentUser = data.user;
            showAppSection();
            await loadCategories();
            await loadTodos();
            
            // Clear form
            document.getElementById('register-form').reset();
        } else {
            const error = await response.text();
            document.getElementById('register-error').textContent = error || 'Registration failed';
        }
    } catch (error) {
        console.error('Registration error:', error);
        document.getElementById('register-error').textContent = 'Network error. Please try again.';
    }
}

function logout() {
    localStorage.removeItem('todo_token');
    currentUser = null;
    categories = [];
    todos = [];
    selectedCategoryFilter = null;
    showAuthSection();
}

// ===================================================================
// CATEGORY FUNCTIONS
// ===================================================================

async function loadCategories() {
    try {
        const response = await apiCall('/categories');
        if (response.ok) {
            categories = await response.json();
            updateCategoryUI();
        } else {
            console.error('Error loading categories:', response.statusText);
        }
    } catch (error) {
        console.error('Error loading categories:', error);
    }
}

function updateCategoryUI() {
    // Update category dropdown for new todos
    const todoSelect = document.getElementById('todo-category');
    todoSelect.innerHTML = '<option value="">No Category</option>';
    categories.forEach(cat => {
        const option = document.createElement('option');
        option.value = cat.id;
        option.textContent = cat.name;
        todoSelect.appendChild(option);
    });
    
    // Update category filters
    const filtersDiv = document.getElementById('category-filters');
    filtersDiv.innerHTML = '';
    categories.forEach(cat => {
        const div = document.createElement('div');
        div.className = 'category-item';
        div.onclick = () => filterByCategory(cat.id);
        div.innerHTML = `
            <div class="category-color" style="background-color: ${cat.color};"></div>
            <span>${cat.name}</span>
        `;
        filtersDiv.appendChild(div);
    });
    
    // Update categories management list
    const categoriesList = document.getElementById('categories-list');
    categoriesList.innerHTML = '<h4>Existing Categories</h4>';
    categories.forEach(cat => {
        const div = document.createElement('div');
        div.className = 'category-item';
        div.innerHTML = `
            <div class="category-color" style="background-color: ${cat.color};"></div>
            <span>${cat.name}</span>
            <button class="btn-danger" onclick="deleteCategory(${cat.id})" 
                    style="margin-left: auto; padding: 4px 8px; font-size: 12px;">Delete</button>
        `;
        categoriesList.appendChild(div);
    });
}

async function addCategory() {
    const name = document.getElementById('category-name').value.trim();
    const description = document.getElementById('category-description').value.trim();
    const color = document.getElementById('category-color').value;
    
    if (!name) {
        alert('Category name is required!');
        return;
    }
    
    try {
        const response = await apiCall('/categories', 'POST', {
            name, description, color
        });
        
        if (response.ok) {
            document.getElementById('category-name').value = '';
            document.getElementById('category-description').value = '';
            document.getElementById('category-color').value = '#3B82F6';
            await loadCategories();
        } else {
            const error = await response.text();
            alert(`Error adding category: ${error}`);
        }
    } catch (error) {
        console.error('Error adding category:', error);
        alert('Error adding category');
    }
}

async function deleteCategory(id) {
    if (!confirm('Are you sure? This will remove the category from all todos.')) {
        return;
    }
    
    try {
        const response = await apiCall(`/categories/${id}`, 'DELETE');
        
        if (response.ok) {
            await loadCategories();
            await loadTodos(); // Refresh todos to update display
        } else {
            const error = await response.text();
            alert(`Error deleting category: ${error}`);
        }
    } catch (error) {
        console.error('Error deleting category:', error);
        alert('Error deleting category');
    }
}

function filterByCategory(categoryId) {
    selectedCategoryFilter = categoryId;
    
    // Update selected filter UI
    document.querySelectorAll('.category-item').forEach(item => {
        item.classList.remove('selected');
    });
    
    if (categoryId === null) {
        document.getElementById('filter-all').classList.add('selected');
    } else {
        // Find and select the clicked category
        const categoryItems = document.querySelectorAll('.category-item');
        categoryItems.forEach(item => {
            if (item.onclick && item.onclick.toString().includes(`filterByCategory(${categoryId})`)) {
                item.classList.add('selected');
            }
        });
    }
    
    displayTodos();
}

// ===================================================================
// TODO FUNCTIONS
// ===================================================================

async function loadTodos() {
    try {
        const response = await apiCall('/todos');
        if (response.ok) {
            todos = await response.json();
            displayTodos();
        } else {
            console.error('Error loading todos:', response.statusText);
        }
    } catch (error) {
        console.error('Error loading todos:', error);
    }
}

function displayTodos() {
    const container = document.getElementById('todos');
    
    // Filter todos based on selected category
    let filteredTodos = todos;
    if (selectedCategoryFilter !== null) {
        filteredTodos = todos.filter(todo => 
            todo.category_id === selectedCategoryFilter
        );
    }
    
    const categoryName = selectedCategoryFilter === null 
        ? 'All Categories'
        : categories.find(c => c.id === selectedCategoryFilter)?.name || 'Unknown';
    
    container.innerHTML = `<h3>Todos - ${categoryName} (${filteredTodos.length})</h3>`;
    
    if (filteredTodos.length === 0) {
        container.innerHTML += '<p style="text-align: center; color: #666; padding: 40px;">No todos in this category yet. Add one above!</p>';
        return;
    }
    
    filteredTodos.forEach(todo => {
        const todoDiv = document.createElement('div');
        todoDiv.className = `todo ${todo.completed ? 'completed' : ''}`;
        
        const categoryBadge = todo.category ? 
            `<div class="category-badge" style="background-color: ${todo.category.color};">
                ${todo.category.name}
            </div>` : '';
        
        todoDiv.innerHTML = `
            ${categoryBadge}
            <div class="todo-title">${escapeHtml(todo.title)}</div>
            <div class="todo-description">${escapeHtml(todo.description) || 'No description'}</div>
            <div class="todo-meta">Created: ${new Date(todo.created_at).toLocaleString()}</div>
            <div class="todo-actions">
                <button class="btn-success" onclick="toggleComplete(${todo.id}, ${!todo.completed})">
                    ${todo.completed ? 'Mark Incomplete' : 'Mark Complete'}
                </button>
                <button class="btn-secondary" onclick="editTodo(${todo.id})">Edit</button>
                <button class="btn-danger" onclick="deleteTodo(${todo.id})">Delete</button>
            </div>
        `;
        
        container.appendChild(todoDiv);
    });
}

async function addTodo() {
    const title = document.getElementById('todo-title').value.trim();
    const description = document.getElementById('todo-description').value.trim();
    const categoryId = document.getElementById('todo-category').value;
    
    if (!title) {
        alert('Title is required!');
        return;
    }
    
    const payload = { title, description };
    if (categoryId) {
        payload.category_id = parseInt(categoryId);
    }
    
    try {
        const response = await apiCall('/todos', 'POST', payload);
        
        if (response.ok) {
            document.getElementById('todo-title').value = '';
            document.getElementById('todo-description').value = '';
            document.getElementById('todo-category').value = '';
            await loadTodos();
        } else {
            const error = await response.text();
            alert(`Error adding todo: ${error}`);
        }
    } catch (error) {
        console.error('Error adding todo:', error);
        alert('Error adding todo');
    }
}

async function toggleComplete(id, completed) {
    try {
        const response = await apiCall(`/todos/${id}`, 'PUT', { completed });
        
        if (response.ok) {
            await loadTodos();
        } else {
            const error = await response.text();
            alert(`Error updating todo: ${error}`);
        }
    } catch (error) {
        console.error('Error updating todo:', error);
        alert('Error updating todo');
    }
}

async function editTodo(id) {
    const todo = todos.find(t => t.id === id);
    if (!todo) return;
    
    const newTitle = prompt('Edit title:', todo.title);
    if (newTitle === null) return; // User cancelled
    
    const newDescription = prompt('Edit description:', todo.description || '');
    if (newDescription === null) return; // User cancelled
    
    try {
        const response = await apiCall(`/todos/${id}`, 'PUT', {
            title: newTitle.trim() || todo.title,
            description: newDescription.trim()
        });
        
        if (response.ok) {
            await loadTodos();
        } else {
            const error = await response.text();
            alert(`Error updating todo: ${error}`);
        }
    } catch (error) {
        console.error('Error updating todo:', error);
        alert('Error updating todo');
    }
}

async function deleteTodo(id) {
    if (!confirm('Are you sure you want to delete this todo?')) {
        return;
    }
    
    try {
        const response = await apiCall(`/todos/${id}`, 'DELETE');
        
        if (response.ok) {
            await loadTodos();
        } else {
            const error = await response.text();
            alert(`Error deleting todo: ${error}`);
        }
    } catch (error) {
        console.error('Error deleting todo:', error);
        alert('Error deleting todo');
    }
}

// ===================================================================
// UTILITY FUNCTIONS
// ===================================================================

function escapeHtml(text) {
    if (!text) return '';
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}