// 
function main () {
  const inputText = document.getElementById("input")
  const submitBtn = document.getElementById("submit")
  submitBtn.addEventListener("click", (e) => {
    e.preventDefault()
    const request = new XMLHttpRequest();
    request.open("POST", "/run", true);
    request.responseType = "text"
    request.send(inputText.value)
    request.onload = () => {
      const stdOutText = document.getElementById("stdout")
      stdOutText.value = ` ${request.response}`
    }
  })
}

window.addEventListener("load", main)
