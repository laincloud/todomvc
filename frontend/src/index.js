import 'todomvc-app-css/index.css';
import 'todomvc-common/base.css';
import 'todomvc-common/base.js';
import Vue from 'vue';
import VueResource from 'vue-resource';

Vue.use(VueResource);

var filters = {
    all: function (todos) {
        return todos;
    },
    active: function (todos) {
        return todos.filter(function (todo) {
            return !todo.completed;
        });
    },
    completed: function (todos) {
        return todos.filter(function (todo) {
            return todo.completed;
        });
    }
};

var app = new Vue({
    data: {
        todos: [],
        newTodo: '',
        editedTodo: null,
        visibility: 'all'
    },
    computed: {
        filteredTodos: function () {
            return filters[this.visibility](this.todos);
        },
        remaining: function () {
            return filters.active(this.todos).length;
        },
        allDone: {
            get: function () {
                return this.remaining === 0;
            },
            set: function (value) {
                this.todos.forEach(function (todo) {
                    todo.completed = value;
                });
            }
        }
    },
    methods: {
        pluralize: function (word, count) {
            return word + (count === 1 ? '' : 's');
        },
        fetchTodos: function () {
            this.$http.get('/todos', {
                headers: {
                    'Accept': 'application/vnd.laincloud.todomvc.v1+json'
                }
            }).then(response => {
                if (response.status !== 200) {
                    console.error('GET /todos failed, response: %o.', response);
                    return;
                }

                console.info('GET /todos succeed, response: %o.', response);
                this.todos = JSON.parse(response.body);
                console.info('this.todos: %o', this.todos);
            }, response => {
                console.error('GET /todos failed, error: %o.', response);
            });
        },
        addTodo: function () {
            var title = this.newTodo && this.newTodo.trim();
            if (!title) {
                return;
            }

            var todo = {
                title: title,
                completed: false
            };
            this.$http.post('/todos', JSON.stringify(todo), {
                headers: {
                    'Accept': 'application/vnd.laincloud.todomvc.v1+json',
                    'Content-Type': 'application/vnd.laincloud.todomvc.v1+json'
                }
            }).then(response => {
                if (response.status !== 201) {
                    console.error('POST /todos failed, body: %o, response: %o.', todo, response);
                    return;
                };

                console.info('POST /todos succeed, body: %o, response: %o.', todo, response);
                var newTodo = JSON.parse(response.body);
                this.todos.push(newTodo);
                this.newTodo = '';
                console.info('this.todos: %o', this.todos);
            }, error => {
                console.error('POST /todos failed, body: %o, error: %o.', newTodo, error);
            });
        },
        removeTodo: function (todo) {
            var index = this.todos.indexOf(todo);
            this.todos.splice(index, 1);
            this.$http.delete('/todos/' + todo.id, {
                method: 'DELETE',
                headers: {
                    'Accept': 'application/vnd.laincloud.todomvc.v1+json'
                }
            }).then(response => {
                if (response.status !== 204) {
                    console.error('DELETE /todos/%s failed, response: %o.', todo.id, response);
                    return;
                }

                console.info('DELETE /todos/%s succeed.', todo.id);
            }, error => {
                console.error('DELETE /todos/%s failed, error: %o.', todo.id, error);
            });
        },
        editTodo: function (todo) {
            this.beforeEditCache = todo.title;
            this.editedTodo = todo;
        },
        doneEdit: function (todo) {
            if (!this.editedTodo) {
                return;
            }
            this.editedTodo = null;
            todo.title = todo.title.trim();
            if (!todo.title) {
                this.removeTodo(todo);
                return;
            }
            this.$http.put('/todos' + todo.id, JSON.stringify(todo), {
                headers: {
                    'Accept': 'application/vnd.laincloud.todomvc.v1+json',
                    'Content-Type': 'application/vnd.laincloud.todomvc.v1+json'
                }
            }).then(response => {
                if (response.status !== 204) {
                    console.error('PUT /todos/%s failed, body: %o, response: %o.', todo.id, todo, response);
                    return;
                };

                console.info('PUT /todos/%s succeed, body: %o, response: %o.', todo.id, todo, response);
            }, error => {
                console.error('PUT /todos/%s failed, body: %o, error: %o.', todo.id, todo, error);
            });
        },
        cancelEdit: function (todo) {
            this.editedTodo = null;
            todo.title = this.beforeEditCache;
        },
        removeCompleted: function () {
            this.todos = filters.active(this.todos);
        }
    },
    mounted: function () {
        this.fetchTodos();
    },
    directives: {
        'todo-focus': function (el, binding) {
            if (binding.value) {
                el.focus();
            }
        }
    }
});

function onHashChange () {
    var visibility = window.location.hash.replace(/#\/?/, '');
    if (filters[visibility]) {
        app.visibility = visibility;
    } else {
        window.location.hash = '';
        app.visibility = 'all';
    }
}

window.addEventListener('hashchange', onHashChange);
onHashChange();

app.$mount('.todoapp');
