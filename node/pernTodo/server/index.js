const express = require("express");
const app = express();
const cors = require("cors")
const pool = require("./db");

// middle ware
app.use(cors());
app.use(express.json())

//ROUTES//

//create a todo
app.post("/todos", async(req, res) => {
    try {
        const { description } = req.body;
        console.log(description)
        const newTodo = await pool.query(
            "INSERT INTO todo (description) VALUES($1) RETURNING *", 
            [description]
        );
        res.json(newTodo.rows[0])
    } catch (err) {
        console.error(err.message);
    }
});

//get all todos

app.get("/todos", async(req, res) => {
    try {
        const allTodos = await pool.query("SELECT * FROM todo");
        res.json(allTodos.rows);
    } catch (err) {
        console.log(err.message);
    }
})

app.get("/todos/:id", async (req, res) => {
    try {
        const { id } = req.params;
        const todo = await pool.query(
            "SELECT * FROM todo WHERE todo_id = $1", 
            [id]
        );

        res.json(todo.rows[0])
    } catch (err) {
        console.log(err.message);
    }
});

// update a todo

app.put("/todos/:id", async (req, res) => {
    try {
        const { id } = req.params;
        const { description } = req.body;

        console.log(id);
        console.log(description);

        const updateTodo = await pool.query(
            "UPDATE todo SET description = $1 WHERE todo_id = $2",
            [description, id]
        )
        res.json("Todo was updated");
    } catch (err) {
        console.error(err.message);
    }
})

// delete a todo

app.delete("/todos/:id", async (req, res) => {
    try {
        const { id } = req.params;
        const deleteTodo = await pool.query(
            "DELETE FROM todo WHERE todo_id = $1",
            [id]
        );
        res.json("deleted")
    } catch (err) {
        console.error(err.message)
    }
})


app.listen(5001, () => {
    console.log("the server has started")
});
