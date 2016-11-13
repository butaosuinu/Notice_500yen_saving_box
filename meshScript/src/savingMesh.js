// Executeに記述

var url = 'http://[your domain]/api/v1/saving'

var date = new Date()
var savingTime = [
					date.getFullYear(),
					date.getMonth() + 1,
					date.getDate()
				 ].join('/') + ' ' 
				 + date.toLocaleTimeString()

var data = {
	"time": savingTime
}

ajax({
	url :  url,
	data : data,
	type : 'post',
	timeout : 5000,
	success : function (contents) {
		callbackSuccess( {
			resultType : "continue",
			runtimeValues : runtimeValues
		} );
	},
	error : function (request, errorMessage) {
		log("Network error");
		runtimeValues.outputIndex = -1;
		callbackSuccess( {
			resultType : "continue",
			runtimeValues : runtimeValues
		} );
	}
})

return {
	resultType : "pause"
};
