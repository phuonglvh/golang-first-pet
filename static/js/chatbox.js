$(document).ready(function () {
  $('#action_menu_btn').click(function () {
    $('.action_menu').toggle()
  })
})

const SenderType = {
  theirs: '',
  mine: '_send',
}

$(function () {
  const baseURI = '/chat/rooms/'
  const segments = window.location.pathname.split('/')
  if (segments.length < 1)
    return
  const roomId = segments[segments.length - 1]
  const userId = window.sessionStorage.getItem('clientId');

  $('#room_id').text(`Group: ${roomId}`)
  $('#contacts').append(contact(`Group: ${roomId}`, true))

  const uri = `${baseURI}/${roomId}/messages`
  $.get(uri, (data, status) => {
    const messages = JSON.parse(data);
    const items = []
    for (const msg of messages) {
      const item = generateMessage(
        msg.Sender === userId,
        'http://www.gravatar.com/avatar',
        msg.content,
        new Date(parseInt(msg.Timestamp)).toString()
      )
      items.push(item);
    }
    $('#room_total_messages').text(`${items.length} messages`)
    $('#msg_container').append(items)
  });

  var socket = null
  var msgBox = $('#message_input')
  $('#btn_send_message').click(function () {
    const content = msgBox.val()
    if (!content) return false
    if (!socket) {
      alert('Error: There is no socket connection.')
      return false
    }

    socket.send(content)
    msgBox.val('')
    const item = generateMessage(
      true,
      'http://www.gravatar.com/avatar',
      content,
      new Date().toISOString()
    )
    console.log({ item })
    $('#msg_container').append(item)
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
      const item = generateMessage(
        msg.Sender === userId,
        'http://www.gravatar.com/avatar',
        msg.content,
        new Date(parseInt(msg.Timestamp)).toString()
      )
      $('#msg_container').append(item)

    }
  }
})

const generateMessage = (isMine, avatarUrl, content, time) => {
  return isMine ? myMessage(avatarUrl, content, time) : theirMessage(avatarUrl, content, time)
}

const theirMessage = (avatarUrl, content, time) => {
  return (
    `<div class='d-flex justify-content-start mb-4'>
      <div class='img_cont_msg'>
        <img src='${avatarUrl}' class='rounded-circle user_img_msg'>
      </div>
      <div class='msg_cotainer'>
        ${content}
        <!-- <span class='msg_time'>${time}</span> -->
      </div>
    </div>`
  );
}

const myMessage = (avatarUrl, content, time) => {
  return (
    `<div class='d-flex justify-content-end mb-4'>
      <div class='msg_cotainer_send'>
        ${content}
        <!-- <span class='msg_time_send'>${time}</span> -->
      </div>
      <div class='img_cont_msg'>
        <img src='${avatarUrl}' class='rounded-circle user_img_msg'>
      </div>
    </div>`
  );
}

const contact = (roomId, isActive, avatarUrl = 'http://www.gravatar.com/avatar') => {
  return `
    <li ${isActive ? `class='active'` : ''}>
      <div class='d-flex bd-highlight'>
          <div class='img_cont'>
              <img src=${avatarUrl} class='rounded-circle user_img'>
              <span class='online_icon offline'></span>
          </div>
          <div class='user_info'>
              <span>${roomId}</span>
              <!-- <<p>Taherah left 7 mins ago</p> -->
          </div>
      </div>
    </li>`;
}
