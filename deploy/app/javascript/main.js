var ismobile = (/iphone|ipad|ipod|android|blackberry|mini|windows\sce|palm/i.test(navigator.userAgent.toLowerCase()));

var initApp = function() {
	this.actions = {
		toggleMenu: function() {
			var overlay = $('.overlay-hugeinc');
			console.log('here');
			if (overlay.hasClass('open')) {
				overlay.removeClass('open');
				overlay.addClass('open');
			} else if (!overlay.hasClass('close')) {
				overlay.addClass('open');
			}
		}
	}
};

initApp.prototype = {
	constructor:initApp,
	fullPage: function() {
		$('#fullpage').fullpage({
			//Navigation
			menu: true,
			lockAnchors: false,
			navigation: true,
			navigationPosition: 'right',
			showActiveTooltip: false,
			slidesNavigation: true,
			slidesNavPosition: 'bottom',
			sectionSelector: 'section',
			slideSelector: '.slide',
			autoScrolling: false,
			controlArrows: false,
		});
	},
	reveal : function() {
		var Reveal = {
			delay: 200,
			distance: '90px',
			easing: 'ease-in-out',
			scale: 1.1,
			reset: true,
		}
		window.sr = ScrollReveal()
			.reveal( '.light-side .face', $.extend(true, Reveal, {origin: 'left'}) )
			.reveal( '.dark-side .face', $.extend(true, Reveal, {origin: 'right'}) )
	},
	onDomReady: function() {
		var self = this;
		$('#menu-btn').on('click', function() {
			$(this).toggleClass('active');
			self.actions.toggleMenu();
		})
	},
	getApi: function(service, request) {
		if (!request) {
			request = null;
		}
		var self = this;
		$.ajax({
			type: "GET",
			url: "/api/v1/"+ service,
			data: request,
			success: function(data) {
				console.log(data)
			}
		});
	}
};

var init = new initApp();

$(document).ready(function() {
	init.fullPage();
	init.onDomReady();
	init.reveal()

	init.getApi("candidates")
});
