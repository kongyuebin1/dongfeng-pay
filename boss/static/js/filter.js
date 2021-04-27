function filter() {
	$.ajax({
		url: "/filter.html",
		success: function(res) {
			let loc = window.location.toString();
			if (res.Code == 404) {
				if (loc.indexOf("login.html") !== -1) {
					return;
				}
				window.parent.location = "/login.html";
			} else if (res.Code == 200) {
				
				if (loc.indexOf("login.html") !== -1) {
					window.parent.location = "/index.html";
				}
			}
		},

		error: function(e) {
			window.parent.location = "/login.html";
		}
	});
};

$().ready(function() {
	filter();
});