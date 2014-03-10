angular.module('Jutraak').factory('Problems', ['$resource', function ($resource) {
    return $resource('/api/problems/:id', {}, {
        'query': { method: 'GET', isArray: true},
        'get': { method: 'GET'},
        'update': { method: 'PUT'}
    });
}]);

