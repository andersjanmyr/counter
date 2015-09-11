$(function() {
  $('.counter').click(function() {
    $.post('/counter', function(data) {
      $('.counter').text(data);
    });
  });
});
