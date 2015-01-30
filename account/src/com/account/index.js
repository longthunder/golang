var myAppModule = angular.module('myApp', ['ui.bootstrap','ngCookies'])
var envMap = {};
function ConfigController($scope, $http){
  $scope.service="Loading"
  $scope.services=[{Id:"Loading",Name:"Loading..."}]
	load("/services", "services")
	loadFn("/envs", function(data){
    $scope.envs=data;
    angular.forEach($scope.envs, function(value, index){
      envMap[value.Id]=value.Name
    })
  })
  $scope.selectService = function() {
    configs=[];
    loadFn("/services/"+$scope.service+"/serviceconfig", function(data){
      angular.forEach(data, function(value, index){
          var content = JSON.stringify(value.Content,"","\t");
          configs.push({
                  EnvId : value.EnvId,
                  EnvName: envMap[value.EnvId], 
                  Content: content, 
                  InitContent: content
                })
      })
      $scope.configs = configs;
    })
  }

  $scope.save = function() {
     angular.forEach($scope.configs, function(value, index){
        if(value.InitContent != value.Content) {
          updateFn("/services/"+$scope.service+"/envs/" 
              + value.EnvId+"/serviceconfig",value.Content)
        }
     })
  }
//---------------------------------------------------  
  function load(path, propName) {
    loadFn(path, function(data){
      $scope[propName] = data;
    })
  }

  function loadFn(path, fn) {
    var req = {
      method: 'GET',
      url: path,
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

  function postFn(path, body) {
    var req = {
      method: 'POST',
      url: path,
      headers: {
        "Accept":"application/json, text/plain, */*"
      },
      data : body
    }
    $http(req).success(function (data) { 
      alert("Success")   
    }).error(function(e){
      alert(e)
    })
  }
  function updateFn(path, body) {
    var req = {
      method: 'PUT',
      url: path,
      headers: {
        "Accept":"application/json, text/plain, */*"
      },
      data : body
    }
    $http(req).success(function (data) { 
      alert("Success")   
    }).error(function(e, status){
      alert(status)
    })
  }
}
