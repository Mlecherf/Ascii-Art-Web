Description:
    This exercice have for goal. To run a server in golang, for show our previous work on ascii art.
    We created a webpage, where you can tap your text, choice the font-family and see the result.

Author:
    Clément VEY
    Mathis LECHERF

Usage:
    He code work only in localhost. For that you have to run the main2.go in your terminal. "go run main2.go"
    Go on your web browser. In the path field tap. "http://localhost:8080/"
    You are on the website. Now you have two zone.
    On the left the TextZone with the textarea, it's where you tap your text
    On the right you have the font-family Zone, you can choose between 3 font-family "standard / shadow / thinkertoy"
    When you are finish click on the submit button, in the bottom of the page, your text will be display in the font-family choice

Implementation details: algorithm:
    "html/template" It's an implementation we use for "link" our go code and html code. It's use for send value in the HTML code.
	"net/http" it's an implementation for http client and server. We use if for, host our website in local. For test the server for exemple. We can use http.Request for retrive data that was tap or choice in the website.