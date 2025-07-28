# EU AI Code of Practice
Audio format: [https://youtu.be/nJxiahNFKSQ](https://youtu.be/2SCoLm1UQnk)

## How Money Moves Between Banks Using Existing Technology?
- Sending a Digital Message: When someone sends money from their bank account to another person's account at a different bank, their bank doesn't just call the other bank. 
  Instead, they send a highly structured digital message. This message contains all the important details, like who sent the money, who's receiving it, 
  the amount, the currency, and the date.
- The "Languages" (Financial Standards): These digital messages are written in specific "languages" called financial messaging standards. Some of the widely used ones include:
  - SWIFT MT messages: This is a very common and older "language" used by over 11,000 financial institutions globally to exchange millions of messages every day, especially for international payments. It has its own unique way of structuring information, using codes like :32A: for the amount or :50F: for the sender's account.
  - ISO 15022: This standard is mainly used for handling things like buying and selling stocks and bonds across borders.
  - ISO 8583: This is the "language" used for almost all credit and debit card transactions, including those at ATMs.
  - FIX (Financial Information eXchange): This "language" is popular for things like trading stocks and other securities.
  - Proprietary Domestic Standards: Many individual banks or specific countries also use their own unique "languages" for internal operations or for communicating within their borders.
- The Grammar (Syntax) and Meaning (Semantics):
  - Each of these "languages" has its own syntax, which is like its grammar or the physical format of the message. 
    For example, one language might show the currency before the amount, while another shows the amount first.
  - They also have semantics, which is the meaning of the words and terms used in the message. 
    Sometimes, different "languages" might use different words for the exact same concept, like calling the person sending money an "Ordering Customer,
    " a "Payer," or a "Payment Originator".
- The Challenge of Many Languages: Because there are so many different "languages" with their own grammars and vocabularies, 
  it's like trying to have a conversation where everyone speaks a slightly different dialect. 
  This makes it difficult for computers to understand everything automatically without human help to interpret the data. 
  Banks have invested a lot of resources into systems that use these existing standards.
- How Banks Manage (Coexistence and Interoperability): To deal with all these different "languages," banks often use special computer programs called middleware. 
  This middleware acts like a digital translator. It takes information from a message in one "language" (standard) and maps or transforms it into the correct format for another "language". 
  This ability for different standards to work together is known as interoperability. 
  This allows the messages to be processed automatically from start to finish, even if different banks use different systems.

## How ISO 20022 Helps Make That Process Easier or Better

Imagine if, instead of everyone having their own secret code for notes, someone invented a master blueprint or a universal recipe for how any secret note should be made. 
This blueprint wouldn't tell you exactly what words to use, but it would tell you how to define all the important parts of the note, 
like where to put the sender's name, where the amount goes, and what each part means.

ISO 20022 is basically that universal recipe or blueprint for creating financial standards. 
It's a method agreed upon by the financial industry to create consistent message standards across all different banking activities, 
like sending money, buying and selling stocks, or using your credit card.

It works in layers, like building blocks for notes:
- Top layer: Business ideas – This layer figures out the basic ideas, like "who is paying?" (the debtor) or "who is receiving the money?" (the creditor) and 
  what information you need for those actions.
- Middle layer: Logical model – This layer describes all the information needed for a specific action, but without worrying about the exact "language" it's written in. 
  It's like deciding you need a "Name" and an "Address" for someone, no matter if you write it in English, French, or Spanish.
- Bottom layer: The "language" itself (syntax) – This is where you pick the actual "language" (like XML, which uses tags like <Ccy> for currency, or JSON) 
  to physically write the message.

ISO 20022 is gaining a lot of popularity and is set to become the main standard for global cross-border payments by the end of 2025. 
Here’s how it makes things easier and better:
- A Common Dictionary for Banks: ISO 20022 has a special dictionary. This dictionary gives clear, agreed-upon definitions for all the common banking terms and concepts.
  So, even if different banks used to call the person sending money a "Payer," "Payor," or "Payment Originator," with ISO 20022,
  everyone knows they're all talking about the "Debtor" (the party that pays).
  This helps clear up confusion and makes sure computers don't get mixed up.
- Reusable Building Blocks: Think of it like Lego bricks. Once you design a Lego brick for a "door," you can use that same "door" brick in many different Lego houses.
  ISO 20022 does this with its "components" (like "PostalAddress" or "FinancialInstitutionIdentification").
  Since these components are reused across many different types of messages, banks only have to understand and set them up once in their systems.
  This makes it much easier and faster to create new messages or update old ones.
- Better Data Quality: ISO 20022 messages can hold much more information, and it's better organized. This "richer" and "higher-quality" data is great because:
  - Customers can get more details about their payments.
  - Regulators and risk managers can check payments more effectively.
  - Businesses can get better insights from their payment data.
- Smoother "Translations" (Interoperability): Remember the "translators" (middleware)? ISO 20022 acts like a universal language that other languages can easily translate into.
  If a bank sends a message in an older "language," it can be mapped to ISO 20022, and then mapped from ISO 20022 to another "language" needed by the receiving bank.
  This "interoperability" means information can flow seamlessly and automatically ("straight-through processing" or STP) across different systems and even different countries,
  reducing errors and making things more efficient. For example, the securities industry saw a dramatic increase in STP rates, from 60-70% to over 95%,
  after adopting a similar structured standard.
- Flexible "Languages": While ISO 20022 often uses a "language" called XML, it's designed to work with other "languages" too, like JSON.
  This means it can adapt as technology changes, so banks aren't stuck with just one way of communicating.

So, imagine banks used to communicate like different sports teams, each with their own unique hand signals and playbooks. 
ISO 20022 is like introducing a universal rulebook and a common training manual that all teams can use to define their plays and signals. 
It means that even if a team still uses some of their old signals, they can easily "map" them to the universal rulebook, allowing for smoother, more accurate, 
and more detailed communication across all teams, making the whole game (moving money) much more efficient and less prone to errors.

## How does ISO 20022 help with things like preventing fraud or errors?

### Preventing Errors:
- Clearer Instructions (Common Dictionary and Definitions):
  - Remember how we said banks used to have many different words for the same thing, like "Payer," "Payor," or "Payment Originator"?
    ISO 20022 provides a common dictionary with clear, agreed-upon definitions for all these banking terms.
    This means everyone uses the same word for the same concept, reducing confusion and the chance of mistakes that happen when different systems interpret things differently.
  - It's like having a universal instruction manual for every ingredient and step in your recipe.
  - If everyone calls a "cup" a "cup" and knows exactly what that means, you're less likely to add too much or too little of something.
- Smart Checks Before Sending (High-Level Business Validation):
  - ISO 20022 messages are built using something called an XML schema (or JSON schema for APIs). Think of a schema as a super-detailed checklist for your recipe.
    It tells the computer exactly what information needs to be in the message, in what order, and what kind of information it should be (like if a field needs to be a date or a number).
  - Before a message even leaves a bank, these schemas can automatically check if the message follows all the rules.
    This is called "business validation" and it reduces the risk of sending or receiving incorrect data. ]
    If something's wrong, the system can catch it immediately, rather than letting a bad message go out and cause problems later.
  - It's like your recipe telling you, "Hey, you forgot the flour!" or "This isn't a valid cooking temperature!" before you even put it in the oven.
- Automatic Processing (Higher Straight-Through Processing - STP):
  - Because the data is so much clearer and structured with ISO 20022, computers can understand and process messages much better without needing people
    to manually look at them or fix things. This is called "Straight-Through Processing" (STP).
  - When more things are done automatically, there's less chance for humans to make typos or other errors.
    For example, when ISO 15022 (which is similar to ISO 20022 in its data dictionary approach) was adopted in the securities industry,
    the rate of straight-through processing jumped from around 60-70% to over 95%, leading to huge savings by reducing errors and manual work.
  - Imagine if your baking robot could perfectly understand every step of the recipe and do it all on its own,
    instead of you having to read messy notes and accidentally grab salt instead of sugar.

### Preventing Fraud
Richer, Higher-Quality Data for Screening:
- ISO 20022 messages can carry much more information and organize it in a better way compared to older systems.
  This "richer" and "higher-quality" data is super important for fighting fraud.
- Regulators and risk managers (the people whose job it is to prevent bad stuff like money laundering or fraud)
  want better assurance that their screening and testing measures are effective. With more detailed and consistent information in each payment message,
  their computer systems can do a much better job of checking for suspicious activity.
  They can see more details about who is sending money, who is receiving it, and why, making it harder for criminals to hide.
- Think of it like adding more ingredients to your recipe for safety: "source of flour must be organic," "sugar must be fair trade."
  The more precise details you have, the easier it is to ensure everything is legitimate and safe.

In short, ISO 20022 is like giving all the banks a universal, super-detailed, and self-checking recipe book for all financial transactions. 
This means less confusion, fewer mistakes, and more detailed information for everyone, making it much harder for errors to slip through or for bad guys 
to sneak their "ingredients" (money) through the system unnoticed.

## Who Must Comply With The New Standard? 
- All Financial Institutions: By the end of 2025, there's a big plan to help all banks and financial institutions adopt ISO 20022 for all payments and reporting.
  This includes the over 11,000 banks that use the SWIFT platform for payments.
- Market Infrastructures: These are like the big central hubs that help many banks trade and settle things (like stock exchanges or systems that clear payments).
  Many of these market infrastructures (MIs), over 100 of them, are also adopting or planning to adopt ISO 20022.
  Examples include TARGET2-Securities (T2S) in Europe, the US Depository Trust and Clearing Corporation (DTCC), and China Central Securities Depositary and Clearing Co. (CCDC).
- Other Financial Areas: ISO 20022 isn't just for sending simple money payments. It's used across many parts of the financial world, like:
  - Payments (sending money from customers to banks, or bank to bank).
  - Funds (like managing investments).
  - Securities (buying and selling stocks or bonds).
  - Foreign Exchange (Forex) (trading different currencies).
  - Cards (credit and debit card transactions).
  - Trade (things like invoices and guarantees). So, any organization involved in these activities will be affected and encouraged to use it.
- Regulators and Authorities: Even the people who make the rules for how money moves (like central banks or government financial watchdogs) are starting to mandate
  the use of ISO 20022 for regulatory reporting. This means if you need to tell the government about your financial activities, you'll need to use this standard.

So, in short, it's not just banks, but a wide range of financial organizations and systems that are involved in moving money, trading, investing, or reporting to authorities, 
all working towards using this new common "recipe".

## Who Manage ISO20022?
- The Financial Industry Itself: The most important thing is that ISO 20022 is a methodology agreed upon by the financial industry to create consistent message standards.
  It's maintained by a broad community of professionals.
- ISO (International Organization for Standardization): This is the big international group that helps create standards for all sorts of things, not just money.
  They approve ISO 20022 as an official standard.
- SWIFT: You might remember SWIFT from our previous talks. They are a big player in global financial messaging.
  SWIFT has been a major contributor to ISO 20022, even drafting the original design. They also act as the "Registration Authority" (RA) for ISO 20022 under an agreement with ISO. Think of the RA as the guardian of the master recipe book and all its ingredients and definitions.
- Special Groups of Experts: There are different expert groups that help manage the standard:
  - Registration Management Group (RMG): This is the highest group that oversees the whole process of adding new parts to the standard.
    They make sure any new "recipes" are actually needed and will be useful to everyone.
  - Standards Evaluation Groups (SEGs): These are groups of experts for specific areas (like payments, securities, or cards)
    who review new "recipes" to make sure they work well for their specific business.
  - Technical Support Group (TSG): This group helps make sure everyone understands the rules of the "recipe book" correctly.

So, it's a team effort by the global financial industry, with SWIFT playing a very central role as the "librarian" and "guardian" of the standard.

## Who Can Create a New Specification?
The ISO 20022 standard is open, which means any organization can help develop and improve the "recipe book". 
You don't even need to be part of ISO!. In fact, 37 organizations have already started developing new ISO 20022 messages.

If an organization wants to create a new message definition or API resource, here's the step-by-step process:
1. Propose Your Idea (Business Justification):
  - First, you need to write a document called a "business justification". This document explains your idea, what kind of message or information you want to create,
    and why it would be good and useful for others in the financial industry.
  - This idea then goes to the Registration Management Group (RMG) for review and approval. They'll check if there's a real need for it and if it makes sense.
2. Develop the "Recipe" (Candidate Messages):
  - If your idea is approved, you get the "green light" to start working on it. You'll work with the Registration Authority (RA) (which is SWIFT).
  - The RA will give you guidelines and help you design the "logical model" for your message. This is like figuring out all the ingredients and steps for your recipe,
    without worrying how you'll write it down yet.
  - You'll use special tools to design these messages, and you can even reuse existing "Lego bricks" (components) from the ISO 20022 dictionary.
  - Once your message is designed, it's called a "candidate ISO 20022 message". The RA will check it to make sure it follows all the rules.
3. Get it Approved by the Experts (Standards Evaluation):
  - The RA then sends your "candidate message" to the Standards Evaluation Groups (SEGs). These are the groups of experts in specific financial areas (like payments or securities).
  - They'll carefully review your message to make sure it really meets the needs of the people who will use it.
    You'll also need to be part of this review and might need to make changes if they ask.
4. Publish the "Recipe" (Official Registration):
  - Once the SEGs give their approval, your message officially becomes an ISO 20022 compliant message.
  - The RA will then officially add it to the ISO 20022 "repository" (the central collection of all official messages and definitions) and publish it for free on the www.iso20022.org website for everyone to use.
  - Even though it's published, you (the organization that created it) still own the message and will be contacted if someone wants to change it later.

So, it's like inventing a new type of building block for our universal Lego set – you propose it, a committee checks if it's needed and fits, 
then you design it with help from the "Lego master," and finally, experts check if it's strong and useful before it's officially added to the big Lego catalog for everyone to use!

## Learning Resource
- https://www.swift.com/campaign/iso-20022/iso-20022-dummies
