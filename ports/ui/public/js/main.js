
Date.prototype.toJSON = function (key) {
    function f(n) {
        // Format integers to have at least two digits.
        return n < 10 ? '0' + n : n;
    }

    return this.getUTCFullYear()   + '-' +
        f(this.getUTCMonth() + 1) + '-' +
        f(this.getUTCDate())      + 'T' +
        f(this.getUTCHours())     + ':' +
        f(this.getUTCMinutes())   + ':' +
        f(this.getUTCSeconds())   + '.' +
        f(this.getUTCMilliseconds())   + 'Z';
};


angular.module('SharedServices', [])
.config(function ($httpProvider) {
    $httpProvider.responseInterceptors.push('myHttpInterceptor');
    var spinnerFunction = function (data, headersGetter) {
        NProgress.start();
        return data;
    };
    $httpProvider.defaults.transformRequest.push(spinnerFunction);
})
// register the interceptor as a service, intercepts ALL angular ajax http calls
.factory('myHttpInterceptor', function ($q, $window) {
    return function (promise) {
        return promise.then(function (response) {
            // do something on success
            // todo hide the spinner
            //alert('stop spinner');
            NProgress.done();
            return response;

        }, function (response) {
            // do something on error
            // todo hide the spinner
            //alert('stop spinner');
            NProgress.done();
            return $q.reject(response);
        });
    };
});

angular.module('Jutraak', ['ngResource', 'ngRoute', 'ngSanitize', 'SharedServices', 'mgcrea.ngStrap'])

.config(['$routeProvider', function($routeProvider) {
    $routeProvider.
        when('/problems', { templateUrl: 'partials/problem/index.html', controller: 'ProblemIndexCtrl' }).
        otherwise({
            redirectTo: '/problems'
        });
}])

.config(['$datepickerProvider', function($datepickerProvider) {
    angular.extend($datepickerProvider.defaults, {
        dateFormat: 'dd.MM.yyyy',
        autoclose: true,
        startWeek: 1
    });
}]);

