if(document.cookie !== ""){
        location.replace("/profile");
    }


const url = "/login"
let inputForm = document.getElementById("formData")



function ParseResponse(response){
    let doc;
    if (response === "Error") {
        doc = document.getElementById("message")
        const control = doc.parentElement;
        control.className = "alert alert-danger alert-dismissible fade show";

    } else if (response === "Success") {
        location.replace("/profile")
    }
}

inputForm.addEventListener("submit", (e) => {
    e.preventDefault();
    const formdata = new FormData(inputForm)
    console.log(formdata)

    fetch(url,{
        method:"POST",
        body:formdata,
    }).then(
        async response =>  ParseResponse(await response.text())
    )

})

