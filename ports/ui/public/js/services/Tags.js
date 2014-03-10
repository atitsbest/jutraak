angular.module('Jutraak').factory('Tags', ['$http', function ($http) {
    return {
        'query': function() {
            return $http.get('/api/problems/tags');
        }
    };
}]);


