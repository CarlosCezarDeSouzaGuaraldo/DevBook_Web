# DevBook_Web
 
DevBook is a social network developed in the GO programming language, where users can create their profile, make posts and follow other users. These three use cases form the core of the application, while DevBook_Web serves as the client-side application responsible for rendering and making requests to the [DevBook_App](https://github.com/CarlosCezarDeSouzaGuaraldo/DevBook_App) API.

## What do you need to install?

1. [Git](https://git-scm.com/) (any version)
2. [GO](https://go.dev/) (any version)
3. [GitHub Desktop](https://desktop.github.com/) (it's not mandatory)

## How to run?

> Firstly, perform the installation and configuration of the repository [DevBook_App](https://github.com/CarlosCezarDeSouzaGuaraldo/DevBook_App) from branch ```main```

1. Clone the repository [DevBook_Web](https://github.com/CarlosCezarDeSouzaGuaraldo/DevBook_Web) from branch ```main```
2. Create at the root of the project one file ```.env``` based on the ```template.env``` file and fill the .env file according to the settings you want to execute.
3. Open the command prompt and navigate to the root directory of the application.
4. Execute the command line ```go run web``` and type ```Enter```.
5. If everything went well, a message like this should appear in your command prompt: ```WEB application running on localhost:3000```

Then, you can access the application in a browser based on the information provided in the environment variables **HOST** and **PORT** in your ```.env``` file.

<div align="center">
 <img src="https://user-images.githubusercontent.com/66181262/231664315-5a0dc387-decf-4b04-ba23-89363c112cc0.jpg" />
</div>
