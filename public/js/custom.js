function toggleMenu() {
	var x = document.getElementById("hamburgerPopout");
	var y = document.getElementById("hamburgerContainer");
	if (x.style.display === "flex") {
		x.style.display = "none";
		y.dataset.open = "false";
	} else {
		x.style.display = "flex";
		y.dataset.open = "true";
	}
}


function openSponsors(){
	window.location.href= "/sponsors";
}

function createCountDown(){
	var countDownDate = new Date("May 11, 2025 22:00:00").getTime();

	var x = setInterval(function() {

	var now = new Date().getTime();

	var distance = countDownDate - now;

	var days = Math.floor(distance / (1000 * 60 * 60 * 24));
	var hours = Math.floor((distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
	var minutes = Math.floor((distance % (1000 * 60 * 60)) / (1000 * 60));
	var seconds = Math.floor((distance % (1000 * 60)) / 1000);

	document.getElementById("countdown").innerHTML = days + "d " + hours + "h "
	+ minutes + "m " + seconds + "s ";

	if (distance < 0) {
		clearInterval(x);
		document.getElementById("countdown").innerHTML = "EXPIRED";
	}
	}, 1000);
}