---
title: "Ideal PR Message"
date: 2022-10-17T20:46:15+02:00
tags: ["developer", "productivity"]
summary: "How to write commit messages and PR summaries that your team will love"
---

If you work as a software engineer, I would assume that you use git every day. Let me give a quick overview of a work process that I envision. You pick up a story, create a feature branch and start implementing whatever is needed. If the story is small and can be completed in one day, you can end up with a single commit to this branch; if the story is slightly more extensive, you should commit to your repo at least once a day or even more often. After the work is done, you will open a Pull Request to your main (or develop) branch, and someone will need to review it and approve it before it can be merged.

Let's first focus on the first part of that process. For each commit, you will need to write a commit message. It should not just be “stuff”, “testing”, or “yet another one”. Hard, I know. Each commit message should describe what is inside of it and what is the idea behind the changes. Even if in your company you do squash commits all the time - your current branch history is a valuable tool. It can help you get up to speed after a weekend and, more importantly, help PR reviewers see what you were doing. When you have your changes described clearly and concisely, it is much easier for anyone to understand the "why” and the "what” of the change. And it would be possible to switch the view from the overall change of your branch to one specific change of one commit - with explanations of what to expect from it.

Now about PRs themselves. It might look like there is a big difference between PRs assigned to someone from your team, a person you have worked with for years already, and a PR given to someone you have not met. That is not true - I think it is much better to treat all PRs equally. It will make you a much more popular developer and help you write software of a higher quality. Putting some thought into your PR messages and not only a story title with a number will help you save yourself from making unnecessary changes and things that don't make sense.

For me, a good PR summary will have the following:
- Story number and title from the issue tracker.
- Quick overview of changes. It can even be a list of commits, which some tools autogenerate for your PRs. If you followed my previous advice about commit messages - you would not need to edit anything!
- Mention edge cases covered and ones not covered but discussed with the stakeholders. Make it clear what was expected and agreed to be in scope.
- If you have some UI changes, it will be beneficial if you include screenshots or, even better, videos of them.
- Add a section describing how to test your change if it is not apparent. The section can include setup instructions, etc.

Yup, it can look overwhelming, but it is not. Just imagine if you were on the receiving end of this PR - admit it, you would enjoy reviewing it (of course, if the code is good too!).

And one last thing you could consider in the PR comments. When someone asks you to improve something, not only make the change and say “done” in response. You can add gifs!! For our team, it became a good tradition that lightens the mood!

What are your best practices for commit and PR messages? 
