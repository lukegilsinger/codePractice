const API_BASE = '/api';

// Load todos when page loads
document.addEventListener('DOMContentLoaded', loadTodos);

async function loadTodos() {
    try {
        const response = await fetch(`${API_BASE}/todos`);
        const todos = await response.json();
        displayTodos(todos);
    } catch (error) {
        console.error('Error loading todos:', error);
    }
}

function displayTodos(todos) {
    const container = document.getElementById('todos');
    container.innerHTML = '<h3>Your Todos</h3>';
    
    if (todos.length === 0) {
        container.innerHTML += '<p>No todos yet. Add one above!</p>';
        return;
    }
    
    todos.forEach(todo => {
        const todoDiv = document.createElement('div');
        todoDiv.className = `todo ${todo.completed ? 'completed' : ''}`;
        
        todoDiv.innerHTML = `
            <h4>${todo.title}</h4>
            <p>${todo.description || 'No description'}</p>
            <small>Created: ${new Date(todo.created_at).toLocaleString()}</small>
            <div>
                <button class="complete-btn" onclick="toggleComplete(${todo.id}, ${!todo.completed})">
                    ${todo.completed ? 'Mark Incomplete' : 'Mark Complete'}
                </button>
                <button class="delete-btn" onclick="deleteTodo(${todo.id})">Delete</button>
            </div>
        `;
        
        container.appendChild(todoDiv);
    });
}

async function addTodo() {
    const title = document.getElementById('title').value.trim();
    const description = document.getElementById('description').value.trim();
    
    if (!title) {
        alert('Title is required!');
        return;
    }
    
    try {
        const response = await fetch(`${API_BASE}/todos`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ title, description }),
        });
        
        if (response.ok) {
            document.getElementById('title').value = '';
            document.getElementById('description').value = '';
            loadTodos(); // Reload the list
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
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ completed }),
        });
        
        if (response.ok) {
            loadTodos(); // Reload the list
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
            loadTodos(); // Reload the list
        } else {
            alert('Error deleting todo');
        }
    } catch (error) {
        console.error('Error deleting todo:', error);
        alert('Error deleting todo');
    }
}