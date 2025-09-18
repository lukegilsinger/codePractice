const API_BASE = '/api';
let categories = [];
let todos = [];
let selectedCategoryFilter = null;

// Load data when page loads
document.addEventListener('DOMContentLoaded', async () => {
    await loadCategories();
    await loadTodos();
});

// CATEGORY FUNCTIONS
async function loadCategories() {
    try {
        const response = await fetch(`${API_BASE}/categories`);
        categories = await response.json();
        updateCategoryUI();
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
            <button class="btn-danger btn-sm" onclick="deleteCategory(${cat.id})" style="margin-left: auto; padding: 2px 6px; font-size: 11px;">Delete</button>
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
        const response = await fetch(`${API_BASE}/categories`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name, description, color }),
        });
        
        if (response.ok) {
            document.getElementById('category-name').value = '';
            document.getElementById('category-description').value = '';
            document.getElementById('category-color').value = '#3B82F6';
            await loadCategories();
        } else {
            alert('Error adding category');
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
        const response = await fetch(`${API_BASE}/categories/${id}`, {
            method: 'DELETE',
        });
        
        if (response.ok) {
            await loadCategories();
            await loadTodos(); // Refresh todos to update display
        } else {
            alert('Error deleting category');
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
        event.target.closest('.category-item').classList.add('selected');
    }
    
    displayTodos();
}

// TODO FUNCTIONS
async function loadTodos() {
    try {
        const response = await fetch(`${API_BASE}/todos`);
        todos = await response.json();
        displayTodos();
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
    
    container.innerHTML = `<h3>Todos - ${categoryName}</h3>`;
    
    if (filteredTodos.length === 0) {
        container.innerHTML += '<p>No todos in this category yet. Add one above!</p>';
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
            <div class="todo-title">${todo.title}</div>
            <div class="todo-description">${todo.description || 'No description'}</div>
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
        const response = await fetch(`${API_BASE}/todos`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(payload),
        });
        
        if (response.ok) {
            document.getElementById('todo-title').value = '';
            document.getElementById('todo-description').value = '';
            document.getElementById('todo-category').value = '';
            await loadTodos();
        } else {
            alert('Error adding todo');
        }
    } catch (error) {
        console.error('Error adding todo:', error);
        alert('Error adding todo');
    }
}

async function toggleComplete(id, completed) {
    try {
        const response = await fetch(`${API_BASE}/todos/${id}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ completed }),
        });
        
        if (response.ok) {
            await loadTodos();
        } else {
            alert('Error updating todo');
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
        const response = await fetch(`${API_BASE}/todos/${id}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ 
                title: newTitle.trim() || todo.title,
                description: newDescription.trim()
            }),
        });
        
        if (response.ok) {
            await loadTodos();
        } else {
            alert('Error updating todo');
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
        const response = await fetch(`${API_BASE}/todos/${id}`, {
            method: 'DELETE',
        });
        
        if (response.ok) {
            await loadTodos();
        } else {
            alert('Error deleting todo');
        }
    } catch (error) {
        console.error('Error deleting todo:', error);
        alert('Error deleting todo');
    }
}