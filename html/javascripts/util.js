$(document).ready(function(){
	$('.nav-item').on('click', function(){
		$(this).siblings().removeClass('activ');
		$(this).addClass('activ');
	});

	//$('.status').text("Status: Here is a status");
});

$.fn.serializeObject = function()
		{
		    var o = {};
		    var a = this.serializeArray();
		    $.each(a, function() {
		        if (o[this.name] !== undefined) {
		            if (!o[this.name].push) {
		                o[this.name] = [o[this.name]];
		            }
		            o[this.name].push(this.value || '');
		        } else {
		            o[this.name] = this.value || '';
		        }
		    });
		    return o;
		};

function sendRequest(mytype, myurl, mydata){
	var req = $.ajax({
		type: mytype,
		url: myurl,
		data: mydata
	});

	req.done(function(data){
		console.log("Successfully made a "+mytype+" request to "+myurl+" with the payload...");
		console.log(mydata);
		return 1;
	});

	req.fail(function(jqxhr, status){
		console.log("Error sending request: "+status);
		return 0;
	});
}