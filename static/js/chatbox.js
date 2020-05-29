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

  const content = (type, body, time) => {
    return `<div class='msg_cotainer${SenderType[type]}'>
              ${body}
              <span class='msg_time${SenderType[type]}'>${time}</span>
            </div>`;
  }

  const avatar = (imgUrl) => {
    return (
      `<div class='img_cont_msg'>
        <img src='${imgUrl}' class='rounded-circle user_img_msg'>
      </div>`
    );
  }

  const message = (type, body, time, imgUrl) => {
    var statEnd = 'start';
    type === SenderType.mine ? statEnd = 'end' : 'start'
    return `
          <div class='d-flex justify-content-${statEnd} mb-4'>
           ${content(type, body, time)}
            ${avatar(imgUrl)}
          </div>
          `;
  }

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
      // $('#msg_container').append(item)
    }
    console.log({ items})
    $('#room_total_messages').text(`${items.length} messages`)
    $('#msg_container').append(items)
  });
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
