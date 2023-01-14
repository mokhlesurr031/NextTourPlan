<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
    <form action="/tours/post/template" method="post" enctype="application/json">

        <label for="name">Name:</label>
        <input type="text" id="name" name="name">
        <br>

        <label for="pick_up_location">Pickup Location:</label>
        <input type="text" id="pick_up_location" name="pick_up_location">
        <br>

        <label for="description">Description:</label>
        <input type="text" id="description" name="description">
        <br>

        <label for="day_count">Day Count:</label>
        <input type="number" id="day_count" name="day_count">
        <br>

        <label for="cost_per_head">Cost Per Head:</label>
        <input type="number" id="cost_per_head" name="cost_per_head">
        <br>

        <label for="created_by">Created By:</label>
        <input type="number" id="created_by" name="created_by">
        <br>

        <input type="submit" value="Submit">
    </form>
</body>
</html>