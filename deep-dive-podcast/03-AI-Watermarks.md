# AI Watermark
Audio format: [https://youtu.be/C8kJHuNdF2c](https://youtu.be/C8kJHuNdF2c)

## What are the social risks of AI-generated content

AI has gotten so good at making images, sounds, and text that it's often super hard for a person to tell if something was made by a human or a computer. This can lead to some big problems:
- Making it hard to tell what's real: Because AI-generated content can look, sound, or read so real, it's becoming difficult to distinguish it from content made by people. This means you might see a picture or read an article online and not know if it's true or if it was just made up by a computer. This can mess with how much we trust information and each other.
- Deepfakes messing with people's reputations: One of the biggest risks is "deepfakes". These are AI-generated or changed images, sounds, or videos that look exactly like real people, objects, or events, but they're fake and designed to make you think they're true. For example, a deepfake could show a famous person saying something they never actually said. If fake content is wrongly linked to someone, it can really hurt their reputation.
- Spreading fake news and influencing important stuff: AI can be used to create fake news or other made-up content online. This rapid spread of untrue information can cause serious problems for how we get our news, how people decide who to vote for in elections, and overall trust in society.
- Cheating and hiding bad behavior: AI can also be used for cheating on schoolwork, like writing essays or code for students. Beyond school, bad actors might use powerful AI tools to hide plagiarism (copying someone else's work) or create tons of fake online reviews to trick people, making it hard to catch them.
- Making AI itself worse: If too much fake or synthetic (AI-made) content ends up on the internet, it can actually pollute the data that new AI models learn from. Since AI models learn from what they're shown, if they're shown a lot of bad or fake AI-generated content, the quality of future AI models could get worse.
- Difficulty in detection: As AI gets better, it becomes increasingly hard to detect AI-generated content. Current detection methods are vulnerable to attacks that can remove or bypass watermarks, making it tough to prove something is AI-generated. Also, sometimes these detectors can even wrongly accuse human-made text as AI-generated, especially if it's from non-native speakers or people who use writing assistance tools.

## What Are AI Watermark

Imagine you have a special invisible stamp, right? An AI watermark is like that stamp, but for things AI creates – pictures, sounds, and even stories! The idea is that the AI puts this secret mark on its work, and even though you can't usually see or hear it, a special computer program can detect it. This helps us know if something was made by a computer or a human, which helps with all those social risks we talked about earlier.

Here's how these secret stamps can be put on different types of AI-generated content.

### For Images (Pictures)
- Visible Markings: Sometimes, it's pretty simple! The AI might just add a small, visible logo, icon, or text like "AI-Generated" directly onto the picture itself. This is especially common for "deepfakes" (fake images that look very real) to clearly tell you a computer made it. The goal of these visible markings is to inform people directly. But, a downside is that they can be easily removed, or they might make the picture look less good.
- Invisible Watermarks: This is where it gets super sneaky! The AI makes tiny, tiny changes to the picture's pixels (the small dots that make up an image) that are so subtle your eyes can't see them. But, a special computer program knows exactly what to look for and can find these hidden changes to confirm it's AI-made. Imagine it like a secret code embedded within the image data itself.
- Metadata Embeddings: Think of metadata as secret labels or notes attached behind the picture file. These notes can contain information like who created the image, when, or with what tool. For AI-generated images, the AI can add information in these notes saying it was made by AI. However, these notes can sometimes be easily changed or removed, especially when pictures are shared on social media, because those platforms often strip out extra information to save space and protect privacy. To make it stronger, a system called C2PA allows for a cryptographic signature to be added to this metadata, like a special digital seal that shows if the information has been tampered with.
-  Digital Fingerprinting: This isn't exactly a stamp on the image, but more like giving the image a unique "digital ID" based on its patterns. It's like a computer looking at a picture and saying, "Aha! This pattern looks just like patterns usually found in AI-generated images!". This method doesn't change the picture itself, but instead creates a unique code (called a hash) that can be compared to a big database of known AI-generated content. YouTube's Content ID system, which finds copyrighted videos, works in a similar way.

### For Text (Stories and Writing)
- Invisible Watermarks (Green/Red Lists): This is super clever for AI that writes! When an AI writes a story or an essay, it usually picks the next word from a huge list of possible words. An AI watermark for text works by having a secret key that divides all possible words into two groups: "green" words and "red" words. The AI is then programmed to slightly prefer using the "green" words and try to avoid the "red" words, without making the writing sound weird or unnatural.

### For Audio (Sounds)
- Invisible Frequencies: Watermarking audio content is similar to watermarking images because both involve working with patterns that are hard for humans to notice but easy for computers to detect. For audio, watermarks are usually put into frequencies (like very high or very low sounds) that are imperceptible to human ears (below about 20 Hz or above 20,000 Hz)
- Newer Techniques: There are new methods like AudioSeal, which can embed a watermark into audio and detect it even if the audio has been edited. This is important for "deepfakes" of voices or other audio that might be used to mislead people

In a nutshell, AI watermarks are like invisible stamps or secret codes that AI puts on its work. They help us tell what's real from what's AI-made, but just like with any secret code, smart people are always trying to figure them out, or even copy them, to trick others. So, it's a constant game of hide-and-seek to keep our digital world transparent and trustworthy!

## Watermark Regulation

 What are the main rules in the EU AI Act? The EU AI Act, which starts applying from August 1, 2026, says two super important things about AI-generated content:

 1. Invisible "Machine-Readable" Markings
    - AI systems must put a secret, invisible code into all the content they create (pictures, audio, video, even text!).
    - This code is like a digital tag that only other special computer programs can "read" and understand.
    - The goal is to make it easy for other systems to automatically detect if something was made by AI, even if you can't see it yourself.
    - These markings need to be "effective, interoperable, robust and reliable". Think of it like a hidden brand stamp in the digital data itself.
2. Visible labels for Deepfakes
    - If the AI creates "deepfakes" then the AI must show a clear, visible sign that it's artificially made.
    - This could be a small logo, an icon, or text directly on the content, like "AI-Generated".
    - The idea is that if you share this deepfake online, the warning stays with it.

### Who has to follow these rules
- Mainly, the "providers" (the companies that develop and sell AI systems or models) and "deployers" (companies that use those AI systems to deliver things to you, like apps or websites).
- Most of the responsibility falls on the providers.
- Even companies offering open-source AI systems (where the code is freely available) are not exempt from these transparency rules if their systems are used in the EU.
- If companies don't follow these rules, they could face really big fines, like up to 15 million Euros!

### What are the challenges with making these rules work?
- Not many AI companies are doing it yet: Even though the rules are coming, only a small number of AI image generators currently add proper watermarks (about 38%), and even fewer add visible deepfake labels (only 8%).
- Companies have conflicting goals: AI companies often want to give users the freedom to create content without any visible or invisible signs of AI, which can make it harder to implement watermarks.
- Watermarks can be removed or faked: Some invisible watermarks, especially those based on "metadata" (hidden notes behind the file), can be easily removed or changed.
- Hard for smaller companies: Forcing companies to add visible labels only for deepfakes means they'd need smart AI to figure out what a "deepfake" prompt is, which can be hard and expensive for smaller businesses. So, some just put visible marks on all content, or offer a paid option to remove them, which probably won't meet the rules.
- Global reach vs. local rules: Many popular AI tools are made by companies outside the EU, so enforcing these rules across the world can be tricky.

## Can Watermarks Be Removed?
The answer is yes, watermarks, especially the "invisible" kind, can be removed or even copied, which is a major concern.
- Visible watermarks are easy to get rid of. If a watermark is a visible logo or text on an image, it can often be easily cropped out or edited away.
- Invisible watermarks (like metadata) can be stripped away. Sometimes it's even automatically removed when you share content on social media (like to save space). Even combining metadata with strong security like a "cryptographic signature" (which proves it's real) or "digital fingerprinting" (a unique code for the content) still relies on many different companies agreeing to keep that information.
- The scary part: "Watermark Stealing" for AI-generated text. This is a clever and serious trick that can make watermarks almost useless.
  - What it is: Imagine someone figures out the secret rules your AI uses to put its invisible stamp on text. They do this by asking the AI's "API" (like its special doorway) to generate a lot of text, and by studying the hidden patterns, they can learn the rules of the secret stamp.
  - It's cheap and works well! An attacker can do this for very little money (under $50) for advanced text watermarks. Once they know the rules, they can use them with a very high success rate (over 80%).
  - Two big dangers once rules are stolen:
    - Spoofing: The attacker can create their own text (even harmful or fake text) and make it look like it has the AI's official stamp, even though it didn't come from that AI. This could make it seem like the AI company (or its owner) is responsible for bad content.
    - Scrubbing: The attacker can take text that was made by an AI and remove the watermark, making it look like a human wrote it. This helps people hide that they used AI for things like cheating on school assignments or creating fake reviews. This is harder to do for very long texts, but watermark stealing makes it much easier.
  - Even "safe" watermarks are at risk: Even advanced text watermarking methods that were thought to be safe against these types of attacks are actually vulnerable.

So, while watermarking is a really important idea to help us know what's real and what's AI-generated, it's like a constant game of digital hide-and-seek and trickery. The people trying to hide AI's involvement are always trying to find ways to remove or fake the "stamps," making it very hard to have a perfect solution right now

##  what is the future of watermarking

- It's an Uphill Battle (but there's hope!)
  - It will be really challenging to get proper watermarks on all AI-generated images, especially because many powerful AI tools are free and "open-source" (meaning anyone can use and change them). As AI models get smaller, it will also become easier for people to make AI images right from their own homes.
  - However, it's good news that the biggest and most powerful AI companies are already starting to put protections in place.
  - Experts agree that more work is needed to create "truly robust" watermarking methods that are super hard to mess with. They also say we need to think about "watermark stealing" (where someone figures out how to copy or remove the stamp) as a major problem to solve.
- Who Should Put the Stamps On?
  - The best way to make watermarks really strong and hard to remove would be for the companies that develop the main AI models (like the "brain" of the AI) to embed the watermarks right from the start. This is called the "model development stage".
  - If the original model creators put the watermarks in, it would be much easier for other companies who use those models in their apps or websites to follow the rules.
  - Governments could also classify the most advanced AI models as "General-Purpose AI models with systemic risks" under the AI Act. This would force their developers to take strong steps, like using robust watermarks, because of the big impact these models can have.
- Smarter Rules and Checks:
  - Since there are so many AI systems, and many are from outside Europe, it will be super important to have automated ways to check if companies are following the rules once the EU AI Act starts. Think of it like a smart computer system automatically checking every digital photo for a hidden code.
  - The rules are not just for images! The EU AI Act also requires watermarks for AI-generated video, audio, and text, and these might be even harder to implement well.
  - Some future watermarking ideas for text might try to be "distortion-free". This means they try to add the watermark without noticeably changing the text, which could make it harder for tricky people to remove them, although this is still a big research question.
- Fighting the Tricky Stuff:
  - One way to make watermarks harder to remove could be to use "multiple secret keys". Imagine having several different secret stamps, making it harder for someone to figure out all the rules to remove them. (Though this still has its own challenges!).
  - AI companies might also make watermarks "selective". This means the watermark (or a stronger version of it) only gets turned on for accounts or content that seem suspicious, like if someone is trying to create fake news.
  - For AI-generated text, one idea to prevent false alarms (where human-written text looks like AI-generated text) is to ignore "repeated n-grams" (which are like common phrases or sequences of words) when checking for a watermark. This would help the watermark focus on the unique "AI-ness" of the text.
  - AI models could even be trained to refuse to do "generative attacks". This means if someone tries to trick the AI into adding or removing watermarks by changing its "language" (like adding an emoji after every word), the AI would simply say no, like a well-trained guard dog.

## Learning Resource
- [Adoption of Watermarking Measures for AI-Generated content and Implications under the EU AI Act](https://arxiv.org/abs/2503.18156)
- [AI Watermarking 101: Tools and Techniques](https://github.com/huggingface/blog/blob/c9293f9d3877baa6f16c20dc243f312ff58eb724/watermarking.md)
- [Watermark Stealing in Large Language Models](https://arxiv.org/abs/2402.19361)
- [A Watermark for Large Language Models](https://arxiv.org/abs/2301.10226)

