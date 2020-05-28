$(function () {
  const baseURI = '/qrcode/generator'
  const location = window.location.pathname
  const roomId = location.match(/\d+/g)
  $('#code').attr('src', `${baseURI}?string=${roomId}`)
})

$(function () {
  var socket = null
  var msgBox = $('#chatbox textarea')
  var messages = $('#messages')

  $('#chatbox').submit(function () {

    if (!msgBox.val()) return false
    if (!socket) {
      alert('Error: There is no socket connection.')
      return false
    }

    socket.send(msgBox.val())
    msgBox.val('')
    return false

  })

  if (!window['WebSocket']) {
    alert('Error: Your browser does not support web sockets.')
  } else {
    var clientId = window.sessionStorage.clientId ? window.sessionStorage.clientId : window.sessionStorage.clientId = Math.floor(Math.random() * 1000000)
    document.cookie = 'X-Authorization=' + clientId + '; path=/'
    socket = new WebSocket(`ws://${window.location.host}` + window.location.pathname + '/ws')
    socket.onclose = function () {
      alert('Connection has been closed.')
    }
    socket.onmessage = (e) => {
      console.log(e)
      messages.append($('<li>').text(e.data))
    }
  }

})
