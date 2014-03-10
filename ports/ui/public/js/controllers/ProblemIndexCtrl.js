angular.module('Jutraak').controller('ProblemIndexCtrl', 
['$scope', '$location', 'Problems', 'Tags', function($scope, $location, Problems, Tags) {
    // Status der Tags.
    $scope.selectedTags = {}

    // Probleme laden.
    $scope.problems = Problems.query();

    // Tags laden.
    $scope.tags = []
    Tags.query().then(function(t) { $scope.tags = t.data });

    
    // Tags aus-/abwählen
    $scope.toggleTag = function(tag) {
        var state = $scope.selectedTags[tag] === true;
        $scope.selectedTags[tag] = !state;
    };

}]);

