var URL = "https://safe-secrets.herokuapp.com/";

async function request(method, url, body = undefined) {
    const res = await fetch('https://safe-secrets.herokuapp.com' + url, { 
      method,
      headers: {
        'Content-Type': 'application/json',
      },
      body: body ? JSON.stringify(body) : undefined,
    });
   
    
    const data = await res.json().catch(e => null);
    if(res.ok) return data;
    
    throw data || res.statusText;
}

async function addSecret(){
    var secret = document.querySelector('#addSecretInput').value;
    var secretUses = document.querySelector('#addSecretDropdown').value;
    var outputKey = document.querySelector("#addSecretOutput");
    var outputID = document.querySelector("#addSecretOutputID");
    var outputRemoval = document.querySelector("#addSecretOutputRemoval");

    var form = document.querySelector(".addSecretForm");
    var outputForm = document.querySelector(".addSecretOutputForm");

    try {
        const data = await postSecret(secretUses, secret);

        form.style="display:none";
        outputForm.style="display:block";

        outputKey.value=data.key;
        outputID.value=data.secret.id;  
        outputRemoval.value=data.removalKey;
    } catch(error) {
        form.style.visibility="hidden";
        outputForm.style.visibility="visible";

        outputKey.value="ERROR: "+error.message;
    }
}

async function deleteSecret(){
    var secretKey = document.querySelector('#deleteSecretInput').value;
    var secretID = document.querySelector('#deleteSecretInputID').value;

    var checkbox = document.querySelector('#deleteSecretCheckbox');
    var status = document.querySelector('#deleteSecretStatus');
    var statusDiv = document.querySelector('.deleteStatusDiv');

    if(checkbox.checked){
        statusDiv.style.visibility = "visible";
        try {
            await request('DELETE',`/?key=${secretID}:${secretKey}`);
            status.value="Secret deleted succesfully";
        } catch(error){
            //handle error
            status.value = "ERROR: "+error.message;
        }
    } else {
        statusDiv.style.visibility = "visible";
        status.value = "Checkbox not checked.";
    }
}

async function getSecret(){
      
    var warning = document.querySelector(".warningDiv");
    var secretKey = document.querySelector('#readSecretInput').value;
    var secretID = document.querySelector('#readSecretInputID').value;
    var output = document.querySelector('#readSecretOutput');

    try {
        const data = await request('GET',`?key=${secretID}:${secretKey}`);
        
        output.value = data.secret.content;

        if(data.secret.usagesLeft!=null) { // null = infinite uses
            if (parseInt(data.secret.usagesLeft) == 0) 
            {
                warning.style.visibility = "visible";
            }
        }
    } catch(error) {
        output.value="ERROR: "+error.message;
    }  
}
async function postSecret(uses,content) {
    var data = {
        "content": content, 
        "uses": parseInt(uses),
    }

    const response = await request('POST','', data);
    return response;
}