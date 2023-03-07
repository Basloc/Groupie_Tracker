var map = L.map('map').setView([51.505, -0.09], 2);


var OpenStreetMap_Mapnik = L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
	maxZoom: 19,
	attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
});


OpenStreetMap_Mapnik.addTo(map)
L.marker("Lyon").addTo(map);


let locations = document.querySelectorAll("li");;
var url = "https://api.opencagedata.com/geocode/v1/json?q="+locations[0]+"&key=ba772045bfb044078998edd6c4dc3c5a";
var request = new XMLHttpRequest();
request.open('GET', url);
request.send()
var data = JSON.parse(request.responseText);
console.log(data.results[0].formatted);