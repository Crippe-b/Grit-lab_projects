// Countdown timer

// 3 minutes from now
var time_in_minutes = 5;
var current_time = Date.parse(new Date());
var deadline = new Date(current_time + time_in_minutes*60*1000);


export function time_remaining(endtime){
	var t = Date.parse(endtime) - Date.parse(new Date());
  var minutes = Math.floor( (t/1000/60) % 60 );
	var seconds = Math.floor( (t/1000) % 60 );
	return {'total':t, 'seconds':seconds + minutes*60};
}

var timeinterval;
export function run_clock(id, endtime){
	function update_clock(){
		var t = time_remaining(endtime);
    document.getElementById("timer").innerText = "TIME: " + t.seconds;
		if(t.total<=0){ clearInterval(timeinterval); }
	}
    update_clock(); // run function once at first to avoid delay
    timeinterval = setInterval(update_clock,1000);
  
}
run_clock('timer', deadline);

var time_left; // time left on the clock when paused

export function pause_clock(){
		clearInterval(timeinterval); // stop the clock
		time_left = time_remaining(deadline).total; // preserve remaining time
	
}

export function resume_clock(){
		// update the deadline to preserve the amount of time remaining
		deadline = new Date(Date.parse(new Date()) + time_left);
		// start the clock
		run_clock('timer', deadline);
}

