var TodoStorage = {
    get: function (todoID) {
        Vue.http.get('/todos/' + todoID).then(function (response) {
            if (response.status !== 200) {
                console.error('GET /todos/%s failed, response: %o.', todoID, response);
                return {};
            }

            console.info('GET /todos/%s succeed, response: %o.', todoID, response);
            return response.json();
        }).catch(function (ex) {
            console.error('GET /todos/%s failed, error: %o.', todoID, ex);
            return {};
        });
    }
};
