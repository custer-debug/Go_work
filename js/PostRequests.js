



    function SendRequest(str){

        var request = new XMLHttpRequest();
        request.open('POST', "/settings/",true);
        console.log(request);
        request.addEventListener('readystatechange', function() {

            if ((request.readyState===4) && (request.status===200)) {
                console.log(request);
                console.log(request.responseText);
            }
        });
        request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
        request.send(str);
        //document.location.replace("/auth/");


    }

        var buttonbar = document.getElementById("send");
        buttonbar.addEventListener("click", SendRequest("delete=true"));


    var but = document.getElementById("Logout");
    but.addEventListener("click", SendRequest("logout=true"));













