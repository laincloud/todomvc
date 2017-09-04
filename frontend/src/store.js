import axios from 'axios';

axios.defaults.headers.common['Accept'] = 'application/vnd.laincloud.todomvc.v1+json';

var TodoStorate = {
    createTodo: function(todos, todo) {
        axios.post('/todos', todo, {
            headers: {
                'Content-Type': 'application/vnd.laincloud.todomvc.v1+json'
            }
        }).then(response => {
            console.info('POST /todos succeed, body: %o, response: %o.', todo, response);
            todos.push(response.data);
        }).catch(error => {
            console.error('POST /todos failed, body: %o, error: %o.', todo, error);
        });
    },
    updateTodo: function(todo) {
        axios.put('/todos/' + todo.id, todo, {
            headers: {
                'Content-Type': 'application/vnd.laincloud.todomvc.v1+json'
            }
        }).then(response => {
            console.info('PUT /todos/%s succeed, body: %o, response: %o.', todo.id, todo, response);
        }).catch(error => {
            console.error('PUT /todos/%s failed, body: %o, error: %o.', todo.id, todo, error);
        });
    },
    deleteTodo: function(todoID) {
        axios.delete('/todos/' + todoID).then(response => {
            console.info('DELETE /todos/%s succeed.', todoID);
        }).catch(error => {
            console.error('DELETE /todos/%s failed, error: %o.', todoID, error);
        });
    }
};

export default TodoStorate;
