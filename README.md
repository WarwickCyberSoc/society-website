# WMG Cyber Society Website

## Contribution
If you'd like to contribute to the website, follow the steps below to get a local copy of it setup.

Install Hugo by following their guide [here.](https://gohugo.io/getting-started/installing)

You'll then want to [fork the git repository](https://docs.github.com/en/github/getting-started-with-github/fork-a-repo#fork-an-example-repository) in order to make a PR later. Once you have forked the repo, clone down your fork.

```
# Note the recurse-submodules, this is required to download the theme.
git clone --recurse-submodules -j8 MYFORKURL
cd society-website
hugo server -D
```

This starts a local version of the website available at `http://localhost:1313`.

Create a new blog post by running the commands below.
```
cd society-website
hugo new post/mypostname/index.md 
cd content/post/mypostname
```

At the top of index.md, you will see the header for the post:

```
---
title: "mypostname"
date: 2021-03-04T10:03:09Z
draft: true
---
```

Firstly add your name as an author parameter.

```
---
title: "Testpost"
date: 2021-03-04T10:03:09Z
author: "John S."
draft: true
---
```

Then you can write your blog post below the header! We use the [Cupper theme](https://github.com/zwbetz-gh/cupper-hugo-theme) which has a [range of shortcodes](https://cupper-hugo-theme.netlify.app/cupper-shortcodes/) which you can use. To add additional assets to your post (such as images), you can place them inside your `mypostname` directory. Then, when referencing them in the blog post, you can use Markdown's standard image tag or Cupper's figure shortcode.

If `image.png` is in the `myblogpost` folder:
```
![Image Alt Text](image.png)

OR

{{< figureCupper img="image.png" caption="This is a caption!" command="Resize" option="700x" >}}
See the shortcode link above for all commands.
```

Your post will be available at `http://localhost:1313/post/myblogpost`. Once you're happy with how it looks, commit your changes to your fork and head to [pull requests](https://github.com/wmgcyber/society-website/compare) on the GitHub repo to make a PR for your fork. Someone will then review your PR and get it pushed to the GitHub repo.
