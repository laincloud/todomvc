import TodoStorage from './store.js';
import './favicon.ico';
import './index.html';

axios.defaults.headers.common['Content-Type'] = 'application/vnd.laincloud.todomvc.v1+json';

var filters = {
    all: function(todos) {
        return todos;
    },
    active: function(todos) {
        return todos.filter(function(todo) {
            return !todo.completed;
        });
    },
    completed: function(todos) {
        return todos.filter(function(todo) {
            return todo.completed;
        });
    }
};

Vue.component('todo-component', {
    props: [
        'id',
        'initialTitle',
        'initialCompleted'
    ],
    template: `
<li class="todo" :class="{completed: completed, editing: editing}">
    <div class="view">
        <input class="toggle" type="checkbox" v-model="completed">
        <label @dblclick="editTodo">{{title}}</label>
        <button class="destroy" @click="removeTodo"></button>
    </div>
    <input class="edit" type="text"
            v-model="title"
            v-todo-focus="editing"
            @blur="doneEdit"
            @keyup.enter="doneEdit"
            @keyup.esc="cancelEdit">
</li>
`,
    data: function() {
        return {
            title: this.initialTitle,
            completed: this.initialCompleted,
            editing: false
        };
    },
    watch: {
        initialCompleted: function(newCompleted, oldCompleted) {
            this.completed = newCompleted;
        },
        completed: function(newCompleted, oldCompleted) {
            if (newCompleted !== oldCompleted) {
                TodoStorage.updateTodo({
                    id: this.id,
                    title: this.title,
                    completed: newCompleted
                });
                this.$emit('update-todo-status', this.id, newCompleted);
            }
        }
    },
    methods: {
        editTodo: function() {
            this.beforeEditCache = this.title;
            this.editing = true;
        },
        doneEdit: function() {
            if (!this.editing) {
                return;
            }

            this.editing = false;
            var title = this.title.trim();
            if (!title) {
                this.removeTodo();
                return;
            }

            TodoStorage.updateTodo({
                id: this.id,
                title: title,
                completed: this.completed
            });
        },
        removeTodo: function() {
            TodoStorage.deleteTodo(this.id);
            this.$emit('remove-todo', this.id);
        },
        cancelEdit: function() {
            this.editing = false;
            this.title = this.beforeEditCache;
        }
    },
    directives: {
        'todo-focus': function(el, binding) {
            if (binding.value) {
                el.focus();
            }
        }
    }
});

var app = new Vue({
    data: {
        todos: [],
        newTodo: '',
        editingTodoId: null,
        visibility: 'all'
    },
    computed: {
        filteredTodos: function() {
            return filters[this.visibility](this.todos);
        },
        remaining: function() {
            return filters.active(this.todos).length;
        },
        allDone: {
            get: function() {
                return this.remaining === 0;
            },
            set: function(value) {
                this.todos.forEach(function(todo) {
                    todo.completed = value;
                });
            }
        }
    },
    methods: {
        pluralize: function(word, count) {
            return word + (count === 1 ? '' : 's');
        },
        addTodo: function() {
            var title = this.newTodo && this.newTodo.trim();
            if (!title) {
                return;
            }

            TodoStorage.createTodo(this.todos, {
                title: title,
                completed: false
            });
            this.newTodo = '';
            console.info('this.todos: %o', this.todos);
        },
        updateTodoStatus: function(todoID, newStatus) {
            this.todos.forEach((item, i, todos) => {
                if (item.id === todoID) {
                    todos[i].completed = newStatus;
                }
            });
        },
        removeTodo: function(todoID) {
            var index = this.todos.findIndex(item => {
                return item.id === todoID;
            });
            this.todos.splice(index, 1);
        },
        removeCompleted: function() {
            var toRemoveTodos = filters.completed(this.todos);
            toRemoveTodos.forEach(todo => {
                TodoStorage.deleteTodo(todo.id);
            });
            this.todos = filters.active(this.todos);
        }
    },
    mounted: function() {
        axios.get('/todos').then(response => {
            console.info('GET /todos succeed, response: %o.', response);
            this.todos = response.data;
        }).catch(error => {
            console.error('GET /todos failed, error: %o.', error);
        });
    }
});

function onHashChange() {
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
