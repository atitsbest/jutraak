angular.module('Jutraak').controller('ProblemIndexCtrl', 
['$scope', '$location', 'Problems', 'Tags', function($scope, $location, Problems, Tags) {

    // Domain-Event: Filter wurde geändert.
    function _emitProblemFilterChangedEvent() {
        $scope.$emit(
                "problemFilterChanged", 
                _.keys($scope.selectedTags),
                $scope.problemQuery);
    }
    //
    // Status der Tags.
    $scope.selectedTags = {}
    // Probleme laden.
    $scope.problems = Problems.query({tags:['t1','t2'], q:'abcsd'});
    
    // Tags laden.
    $scope.tags = []
    Tags.query().then(function(t) { $scope.tags = t.data });

    
    // Tags aus-/abwählen
    $scope.toggleTag = function(tag) {
        if ($scope.selectedTags[tag] !== undefined) {
            delete $scope.selectedTags[tag];
        }
        else {
            $scope.selectedTags[tag] = true;
        }

        // Domain-Event: Filter wurde geändert.
        _emitProblemFilterChangedEvent();
    };

    $scope.$watch("problemQuery", $.debounce(500, _emitProblemFilterChangedEvent));


    // Wenn sich der Filter geändert hat, müssen wir die 
    // Probleme neu laden.
    $scope.$on('problemFilterChanged', function(e, tags, q) {
        $scope.problems = Problems.query({tags:tags, q:q});
    });
}]);

