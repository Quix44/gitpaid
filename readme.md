# GitPaid

### Description

GitPaid is an innovative platform that connects open source projects with developers, streamlining the way developers find work and contribute to projects. Here's how it simplifies the process:

Users log in using their GitHub accounts. This login also automatically creates a digital wallet for them within the app, and lists their GitHub projects on a dedicated page.

If a project owner needs developers, they can link their GitHub project to GitPaid and add funds to it. This is done by copying a specific webhook URL into their project's settings on GitHub and depositing a chosen amount of digital currency (ERC-20 tokens) through GitPaid.

Owners can then mark specific project issues with a bounty (like "100 usdc") using GitHub. These tagged issues appear on GitPaid for all logged-in users, showing the problem to be solved and the bounty offered.

Developers interested in tackling these issues can find them based on programming language or bounty size. To work on an issue, they submit a solution via GitHub. A project maintainer reviews the submission, and if approved, the issue is marked as resolved, automatically paying the developer the bounty.

GitPaid democratizes the process of finding development work. Unlike platforms such as Fiverr or Airtasker, where job opportunities can be limited by biases or the need for positive reviews, GitPaid offers a level playing field. By using blockchain technology, it ensures transparency and equal opportunity for all developers to earn by solving issues.

# How it's Made

Frontend Development: The user interface is all NextJS, incorporating shadcn components.

Cloud Services: Amazon Web Services (AWS) support the app's backend. It has two main functions hosted on AWS Lambda (in Golang): one handles notifications from GitHub, and the other manages requests from the app's interface. AWS's DynamoDB is used to store data.

GitHub Integration: The app is closely connected with GitHub, utilizing GitHub's notification services and user interface to link everything together smoothly.

Smart Contract: The app includes a smart contract, making it deployable across various blockchain networks, we deployed on Base and Arbitrum.

A key technology partner is the 'Dynamic Wallet', which simplifies the complex aspects of blockchain (web3) technology, offering users a straightforward experience similar to traditional web (web2) applications.

How these technologies come together:

Logging In: Users can sign into the app with their GitHub accounts. This process is enhanced by NextJS and the dynamic wallet feature.

Loading Data: After logging in, the app fetches user data using AWS Lambda functions and stores it in DynamoDB.

Interaction and Transactions: Users interact with the app and GitHub, where smart contracts are executed automatically when needed, thanks to the seamless integration of these technologies. This approach allows the app to extend GitHub's capabilities into the blockchain world, making it a unique and creative web3 application.
