import 'directory/build/director.js';
import 'todomvc-app-css/index.css';
import 'todomvc-common/base.css';
import 'todomvc-common/base.js';
import 'vue/dist/vue.js';
import 'whatwg-fetch';

import './route.js';

(function (exports) {
    'use strict';

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

    exports.app = new Vue({
        el: '.todoapp',
        data: {
            todos: todoStorage.fetchList(),
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
            addTodo: function () {
                var value = this.newTodo && this.newTodo.trim();
                if (!value) {
                    return;
                }

                var newTodo = todoStorage.create({
                    title: value,
                    completed: false
                });
                this.todos.push(newTodo);
                this.newTodo = '';
            },
            removeTodo: function (todo) {
                var index = this.todos.indexOf(todo);
                this.todos.splice(index, 1);
                todoStorage.delete(todo.id);
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
                todoStorage.update(todo.id, todo);
            },
            cancelEdit: function (todo) {
                this.editedTodo = null;
                todo.title = this.beforeEditCache;
            },
            removeCompleted: function () {
                this.todos = filters.active(this.todos);
            }
        },
        directives: {
            'todo-focus': function (el, binding) {
                if (binding.value) {
                    el.focus();
                }
            }
        }
    });
})(window);
