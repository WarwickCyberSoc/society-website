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
	window.location.href("/sponsors/");
}