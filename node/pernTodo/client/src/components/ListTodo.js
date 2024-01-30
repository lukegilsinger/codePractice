import React, { Fragment, useEffect, useState } from "react";
import EditTodo from "./EditTodo"

const ListTodo = () => {
    
    const [todos, setTodos] = useState([])

    const getTodos = async () => {
        try {
            const response = await fetch("http://localhost:5001/todos")
            const jsonData = await response.json();

            setTodos(jsonData);
        } catch (err) {
            console.error(err.message);
        }
    }
    useEffect(() => {
        getTodos();
    }, []);

    const deleteTodo = async(id) => {
        try {
            const dt = await fetch(`http://localhost:5001/todos/${id}`, {
                method: "DELETE"
            });
            setTodos(todos.filter(todo => todo.todo_id !== id));
            console.log(dt)
        } catch (err) {
            console.error(err.message)
        }
    }
    return (
<Fragment>
    <table className="table">
        <thead>
            <tr>
                <th>Action</th>
                <th>Lastname</th>
                <th>D</th>
            </tr>
        </thead>
        <tbody>
            {/* <tr>
                <td>John</td>
            </tr> */}
            {todos.map(todo => (
                <tr key={todo.todo_id}>
                    <td>{todo.description}</td>
                    <td>
                        <EditTodo todo={todo}/> 
                    </td>
                    <td>
                        <button className="btn btn-danger"onClick={() => deleteTodo(todo.todo_id)}>Delete</button>
                    </td>
                </tr>
            ))}
        </tbody>
    </table>

</Fragment>
    )
}

export default ListTodo;