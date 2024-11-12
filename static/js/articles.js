document.addEventListener("DOMContentLoaded", () => {
    const articlesContainer = document.querySelector('.articles');

    fetch('/api/articles')
        .then(response => response.json())
        .then(articles => {
            articlesContainer.innerHTML = '';

            articles.forEach(article => {
                const articleElement = document.createElement('article');
                articleElement.classList.add('article');

                articleElement.innerHTML = `
          <h2 class="article-title">
            <a href="#" data-id="${article.id}" class="article-link">${article.title}</a>
          </h2>
          <span class="article-date">Published on ${article.date}</span>
          <p class="article-content">Click to read the full article...</p>
        `;

                articlesContainer.appendChild(articleElement);
            });

            document.querySelectorAll('.article-link').forEach(link => {
                link.addEventListener('click', function(event) {
                    event.preventDefault();
                    const articleId = this.getAttribute('data-id');
                    fetchAndDisplayArticle(articleId);
                });
            });
        })
        .catch(error => console.error('Error fetching articles:', error));
});

function fetchAndDisplayArticle(id) {
    fetch(`/api/article/${id}`)
        .then(response => response.json())
        .then(article => {
            const articlesContainer = document.querySelector('.articles');
            articlesContainer.innerHTML = `
        <article class="article-full">
          <h2 class="article-title">${article.title}</h2>
          <span class="article-date">Published on ${article.date}</span>
          <div class="article-content">${article.content}</div>
          <button onclick="location.reload()">Back to Articles</button>
        </article>
      `;
        })
        .catch(error => console.error('Error fetching article:', error));
}
