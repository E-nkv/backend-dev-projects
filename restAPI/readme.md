Explanation of restAPI:

main.go's only purpose is to get env vars (in this case, they are global), setup dependencies (create an instance of App), mount the multiplexer, and run the server. NOTHING ELSE.

package api takes care of api level interactions: routing, reading and writing from request / response, and calling the corresponding logic in the Service. the overall flow of the handlers are: read relevant data from the request (like url params), call the corresponding service with this data, and then give this response back to the user.

package types is pretty straightforward: it holds all the relevant types.

package service takes care of the heavy logic for each of the endpoints.

This structure, in my experience, works the best, sitting in the sweet spot between simplicity and readability/maintainability:
    - alternative #1 is the "just make it work" approach (learned it from AnthonyGG), which, though simple, becomes a messy mambojambo once the project grows larger, and becomes a pain to maintain / modify.
    - alternative #2 is to use a heavy-weight design pattern (like the repository pattern), having a service which in turns call the repository for all data-related interactions (like talking to the DB). this approach, despite popular, becomes quite tedious and not that practical in real life projects.
