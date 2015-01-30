var myAppModule = angular.module('myApp', ['ui.bootstrap','ngCookies'])

function ConfigController($scope, $http){			
	load("/services", "services")
	load("/envs", "envs")
	/**$scope.sel = function(key, value) {
		$scope[key] = value;
		loadConfig(getParam($scope));
	}
	$scope.mode = function(mode) {
		$scope["mode"] = mode;
	}
  function loadConfig(param) {
    if(param.mode == "service") {
      if(env !=null && service != null) {
        load($scope, $http, "/serviceconfigs/services/"+service, function(list){
          if(list!=null) {
              

          }
        })
      }
    }
  }**/

  function load(path, propName) {
    loadFn(path, function(data){
      $scope[propName] = data;
    })
  }

  function loadFn(path, fn) {
    var req = {
      method: 'GET',
      url: endpoint + path,
      headers: {
        "Accept":"application/json, text/plain, */*"
      }
    }
    $http(req).success(function (data) { 
          fn(data);       
    }).error(function(e){
      alert(e)
    })
  }
}