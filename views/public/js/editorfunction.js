$(function() {
    var $splitter = $('.bar');
    var $left = $('.box-1');
    var $right = $('.box-2');
    var isDragging1 = false;

    // Set the initial width of the left and right elements
    $left.css('flex-basis', '33.00%');
    $right.css('flex-basis', '67.00%');

    // Register the drag event on the first split bar
    $splitter.on('mousedown', function(e) {
      e.preventDefault();
      isDragging1 = true;
    });

    $(document).on('mousemove', function(e) {
      if (isDragging1) {
        window.requestAnimationFrame(function() {
          var totalWidth = $left.width() + $right.width();
          var x = e.pageX - $left.offset().left - $splitter.width() / 2;
          var leftWidth = x / totalWidth * 100 + '%';
          var rightWidth = (totalWidth - x) / totalWidth * 100 + '%';

          // Check to limit the change in width
          if (x >= 0 && x <= totalWidth) {
            $left.css('flex-basis', leftWidth);
            $right.css('flex-basis', rightWidth);
          }
        });
      }
    }).on('mouseup', function(e) {
      isDragging1 = false;
    });
  });

  $(function() {
    var $splitter2 = $('.bar-2');
    var $top = $('.box-2-top');
    var $bottom = $('.box-2-bottom');
    var isDragging2 = false;

    // Set the initial height of the top and bottom elements
    $top.css('flex-basis', '50.00%');
    $bottom.css('flex-basis', '50.00%');

    // Register the drag event on the second split bar
    $splitter2.on('mousedown', function(e) {
      e.preventDefault();
      console.log("pressed")
      isDragging2 = true;
    });

    $(document).on('mousemove', function(e) {
      if (isDragging2) {
        window.requestAnimationFrame(function() {
          var totalHeight = $top.height() + $bottom.height();
          var y = e.pageY - $top.offset().top - $splitter2.height() / 2;
          var topHeight = y / totalHeight * 100 + '%';
          var bottomHeight = (totalHeight - y) / totalHeight * 100 + '%';

          // Check to limit the change in height
          if (y >= 0 && y <= totalHeight) {
            $top.css('flex-basis', topHeight);
            $bottom.css('flex-basis', bottomHeight);
          }
        });
      }
    }).on('mouseup', function(e) {
      isDragging2 = false;
    });
  });
