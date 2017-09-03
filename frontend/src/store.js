(function (exports) {
    'use strict';

    exports.todoStorage = {
        fetchList: function () {
            fetch('/todos').then(function (response) {
                if (response.status != 200) {
                    console.error('GET /todos failed, response: %o.', response);
                    return [];
                }

                console.info('GET /todos succeed, response: %o.', response);
                return response.json();
            }).catch(function (ex) {
                console.error('GET /todos failed, error: %o.', ex);
                return [];
            });
        },
        create: function (todo) {
            fetch('/todos', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/vnd.laincloud.todomvc.v1+json'
                },
                body: JSON.stringify({
                    title: todo.title,
                    done: todo.completed
                })
            }).then(function (response) {
                if (response.status != 201) {
                    console.error('POST /todos failed, body: %o, response: %o.', todo, response);
                    return {};
                };

                console.info('POST /todos succeed, body: %o, response: %o.', todo, response);
                return response.json();
            }).catch(function (ex) {
                console.error('POST /todos failed, body: %o, error: %o.', todo, ex);
                return {};
            });
        },
        get: function (todoID) {
            fetch('/todos/' + todoID).then(function (response) {
                if (response.status != 200) {
                    console.error('GET /todos/%s failed, response: %o.', todoID, response);
                    return {};
                }

                console.info('GET /todos/%s succeed, response: %o.', todoID, response);
                return response.json();
            }).catch(function (ex) {
                console.error('GET /todos/%s failed, error: %o.', todoID, ex);
                return {};
            });
        },
        update: function (todoID, todo) {
            fetch('/todos' + todoID, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/vnd.laincloud.todomvc.v1+json'
                },
                body: JSON.stringify({
                    title: todo.title,
                    done: todo.completed
                })
            }).then(function (response) {
                if (response.status != 204) {
                    console.error('PUT /todos/%s failed, body: %o, response: %o.', todoID, todo, response);
                    return;
                };

                console.info('PUT /todos/%s succeed, body: %o, response: %o.', todoID, todo, response);
            }).catch(function (ex) {
                console.error('PUT /todos/%s failed, body: %o, error: %o.', todoID, todo, ex);
            });
        },
        delete: function (todoID) {
            fetch('/todos/' + todoID, {
                method: 'DELETE'
            }).then(function (response) {
                if (response.status != 204) {
                    console.error('DELETE /todos/%s failed, response: %o.', todoID, response);
                    return;
                }

                console.info('DELETE /todos/%s succeed.', todoID);
            }).catch(function (ex) {
                console.error('DELETE /todos/%s failed, error: %o.', todoID, ex);
            });
        }
    };
})(window);
