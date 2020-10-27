document.addEventListener("DOMContentLoaded", function(){

    let buttonbar = document.getElementById('send');
    buttonbar.addEventListener("click", function(){

        let name = 'delete=true';
        let request = new XMLHttpRequest();
        request.open('POST', "/settings/",true);
              console.log(request);
                              request.addEventListener('readystatechange', function() {
      
      if ((request.readyState===4) && (request.status===200)) {
          console.log(request);
            console.log(request.responseText);
                 
      } 
                                    });
          request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded')
           request.send(name);
        document.location.replace("/auth/");
      });
});