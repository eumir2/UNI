
function controllaForm(e) {
    var email1 = document.getElementById("InputEmail1");
    var email2 = document.getElementById("InputEmail2");
    var password1 = document.getElementById("InputPassword1");
    var password2 = document.getElementById("InputPassword2");
    var nome = document.getElementById("InputNome");
    var cognome = document.getElementById("InputCognome");
    // var codice_fiscale = document.getElementById("InputCF").value
    // var indirizzo = document.getElementById("InputIndirizzo").value
    var genere = document.getElementById("InputGenere");
    console.log(email1);
    console.log(email2);
    console.log(email1 != email2);
    var flag = true
    var errori = ""
    if(controllaLunghezza(email1,6)){
        email1.classList.add("border-danger")
    }
    if (email1.value != email2.value){
        errori += "Email non corrispondenti"
        flag = false;
        email1.classList.add("border-danger")
        email2.classList.add("border-danger")
        //e.style.borderColor = "red"
    }
    if (password1.value != password2.value) {
        errori += "<br>Password non corrispondenti"
        password1.classList.add("border-danger")
        password2.classList.add("border-danger")
        flag = false;
    }
    if (password1.value.length < 8) {
        errori += "<br>Password troppo corta"
        password1.classList.add("border-danger")
        flag = false;
    }

    if(controllaLunghezza(nome,3)){
        flag = false;
        errori += "<br>Nome non compilato"
    }
    if(controllaLunghezza(cognome,5)){
        flag = false;
        errori += "<br>Cognome non compilato"
    }
    if(genere.value == -1){
        flag = false;
        errori += "<br>Genere non selezionato"
        genere.classList.add("border-danger")
    }
    console.log(errori)
    if (!flag) {
        var alertBS = document.getElementById('alert');
        alertBS.innerHTML = errori;
        alertBS.classList.remove("d-none");
        //equivalente
        // document.getElementById('alert').innerHTML = errori;
        // document.getElementById('alert').classList.remove("d-none");
        e.preventDefault();
        return flag;
    }

    function controllaLunghezza(elemento, lunghezza){
        if(elemento.length < lunghezza){
            elemento.classList.add("border-danger")
            return true
        }
        return false

        //return elemento.lenght < lunghezza
    }
  
    alert(email1 + "\n" + email2 + "\n" + password1 + "\n" + password2
        + "\n" + nome + "\n" + cognome + "\n" + codice_fiscale + "\n" + indirizzo + "\n" + genere);
}