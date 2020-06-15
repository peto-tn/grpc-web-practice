<template>
  <div class="articles">
    <div v-for="post in feed" :key="post.articleId">
      {{ post.title }}
    </div>
  </div>
</template>

<script>
const { GetArticleRequest, ListArticleRequest } = require('../proto/api/article_pb')
const { ArticleClient } = require('../proto/api/article_grpc_web_pb')

export default {
  name: 'app',
  components: {},
  data: function () {
    return {
      testData: '',
      posts: [],
    }
  },
  created: function () {
    // eslint-disable-next-line
    this.client = new ArticleClient('http://localhost:8080', null, null)
    this.listArticle()
  },

  computed: {
    feed() {
      return this.posts
    }
  },

  methods: {
    getArticle: function () {
      // eslint-disable-next-line
      let getRequest = new GetArticleRequest()
      // eslint-disable-next-line
      this.client.get(getRequest, {}, (err, response) => {
        // eslint-disable-next-line
        let articleData = response.toObject()
        this.testData = articleData.title
        console.log(articleData.title)
      })
    },
    listArticle: function () {
      // eslint-disable-next-line
      let request = new ListArticleRequest()
      //request.setNumber(Number(this.inputField))
      // eslint-disable-next-line
      this.client.list(request, {}, (err, response) => {
        // eslint-disable-next-line
        console.log(err)
        console.log(response.toObject())
        let articleDataList = response.toObject().articleDataList
        this.posts = articleDataList
        articleDataList.forEach(articleData => {
          this.testData = articleData.title
          console.log(articleData.title);
        })
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
