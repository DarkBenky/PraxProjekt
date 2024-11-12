<template>
    <div class="post">
        <!-- Post Header with User Info -->
        <div class="post-header">
            <UserProfile :user="user" />
            <small class="post-date">{{ formatDate(post.created_at) }}</small>
        </div>

        <!-- Post Content -->
        <div class="post-content" @click="toggleComments">
            <p>{{ post.content_text }}</p>
            <button class="toggle-comments">
                {{ isActive ? 'Hide Comments' : 'Show Comments' }}
            </button>
            <button class="toggle-comments" @click="navigateToPost">
                Navigate to Post
            </button>
            <div class="edit-delete-options" v-if="post.userID === $store.state.userId">
                <button class="toggle-delete" @click="deletePost">
                    Delete
                </button>
                <button class="toggle-edit" @click="editPost">
                    Edit
                </button>
            </div>
        </div>

        <!-- Comments Section -->
        <div v-if="isActive" class="comments-section">
            <div v-if="loadingComments" class="loading">
                Loading comments...
            </div>
            <div v-else-if="commentsError" class="error">
                {{ commentsError }}
            </div>
            <div v-else>
                <div v-if="comments">
                    <!-- Add Comment Section -->
                    <div class="add-comment-section">
                        <textarea v-model="newComment" placeholder="Add a comment..."></textarea>
                        <button @click="addComment">Add Comment</button>
                    </div>

                    <ul v-if="comments.length > 0" class="comments-list">
                        <li v-for="comment in comments" :key="comment.idComment" class="comment">
                            <!-- Comment Header with User Info -->
                            <div class="comment-header">
                                <UserProfile :user="getUserWithId(comment.idUser)" />
                            </div>
                            <div class="comment-content">
                                <p>{{ comment.content_text }}</p>
                                <small class="comment-date">{{ formatDate(comment.created_at) }}</small>
                            </div>
                        </li>
                    </ul>
                </div>
                <p v-else class="no-comments">No comments yet</p>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
import UserProfile from './UserProfile.vue';

export default {
    name: 'PostView',

    components: {
        UserProfile
    },

    props: {
        post: { type: Object, required: true },
        user: { type: Object, required: true },
        users: { type: Array, required: true }
    },

    data() {
        return {
            isActive: false,
            comments: [],
            loadingComments: false,
            commentsError: null,
            baseUrl: "http://localhost:5050",
            newComment: "",
        };
    },

    methods: {
        editPost() {
            this.$router.push(`/edit/${this.post.idPost}`);
        },

        // In Post.vue, update the deletePost method:
        async deletePost(event) {
            event.stopPropagation(); // Prevent event bubbling
            try {
                const response = await axios.delete(`${this.baseUrl}/deletePost`, {
                    data: {
                        postID: String(this.post.idPost)
                    }
                });

                if (response.status === 200) {
                    // Emit event to parent to refresh posts
                    this.$emit('post-deleted', this.post.idPost);
                }
            } catch (error) {
                console.error("Error deleting post:", error);
                alert("Failed to delete post. Please try again.");
            }
        },
        async addComment() {
            // Ensure there is content in the new comment field
            if (!this.newComment.trim()) {
                alert("Comment content cannot be empty.");
                return;
            }

            try {
                // Make a POST request to the backend with the new comment data
                const response = await axios.post(`${this.baseUrl}/addComment`, {
                    postID: String(this.post.idPost),
                    userID: String(this.$store.state.userId),
                    contentText: this.newComment
                });

                // If the request was successful, fetch the updated list of comments
                if (response.status === 200) {
                    this.newComment = ""; // Clear the comment input field
                    await this.fetchComments(this.post.idPost); // Refresh the comments
                }
            } catch (error) {
                alert("Failed to add comment. Please try again.");
                console.error("Error adding comment:", error);
            }
        },
        navigateToPost() {
            this.$router.push(`/post/${this.post.idPost}`);
        },
        getUserWithId(id) {
            if (this.users.length === 0) return null;
            return this.users.find(user => user.idUser === id);
        },
        async toggleComments() {
            if (this.isActive) {
                this.isActive = false;
                this.comments = [];
            } else {
                this.isActive = true;
                await this.fetchComments();
            }
        },

        async fetchComments() {
            this.loadingComments = true;
            this.commentsError = null;

            try {
                const response = await axios.get(`${this.baseUrl}/comments`, {
                    params: {
                        idPost: this.post.idPost
                    }
                });
                this.comments = response.data;
            } catch (error) {
                this.commentsError = 'Failed to load comments: ' + error.message;
                console.error('Error fetching comments:', error);
            } finally {
                this.loadingComments = false;
            }
        },

        formatDate(dateString) {
            try {
                const date = new Date(dateString);
                return date.toLocaleDateString() + ' ' + date.toLocaleTimeString();
            } catch (err) {
                return dateString;
            }
        }
    }
};
</script>

<style scoped>
.add-comment-section {
    display: flex;
    flex-direction: column;
    margin-bottom: 1.5rem;
}

.add-comment-section textarea {
    min-height: 80px;
    padding: 10px;
    margin: 15px;
    border-radius: 4px;
    border: 1px solid #ddd;
    resize: vertical;
    font-size: 1rem;
    font-family: inherit;
    margin-bottom: 0.5rem;
    transition: border-color 0.2s ease-in-out;
}

.add-comment-section textarea:focus {
    border-color: #999;
    outline: none;
    padding: 10px;
    margin: 15px;
}

.add-comment-section button {
    align-self: flex-end;
    padding: 0.5rem 1rem;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 1rem;
    transition: background-color 0.2s ease-in-out;
}

.add-comment-section button:hover {
    background-color: #0056b3;
}

.add-comment-section button:disabled {
    background-color: #cccccc;
    cursor: not-allowed;
}

.post {
    margin-bottom: 1.5em;
    border: 1px solid #ddd;
    border-radius: 8px;
    overflow: hidden;
}

.post-header {
    padding: 0.5em 1em;
    background-color: #fff;
    border-bottom: 1px solid #eee;
}

.post-content {
    padding: 1em;
    cursor: pointer;
    background-color: #f9f9f9;
    transition: background-color 0.2s;
}

.post-content:hover {
    background-color: #f0f0f0;
}

.post-date {
    display: block;
    color: #666;
    margin-top: 0.5em;
}

.toggle-comments {
    margin-top: 0.5em;
    padding: 0.3em 0.8em;
    background-color: #e0e0e0;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

.comments-section {
    padding: 1em;
    background-color: #fff;
    border-top: 1px solid #eee;
}

.comment {
    padding: 0.8em;
    border-bottom: 1px solid #eee;
}

.comment-header {
    margin-bottom: 0.5em;
}

.comment-content {
    padding-left: 1em;
}

.loading,
.error,
.no-comments {
    padding: 1em;
    text-align: center;
}

.error {
    color: #d32f2f;
    background-color: #ffebee;
    border-radius: 4px;
}

.loading {
    color: #666;
    font-style: italic;
}

.no-comments {
    color: #666;
    font-style: italic;
}

.edit-delete-options {
    margin-top: 1rem;
    display: flex;
    gap: 10px;
}

.toggle-delete,
.toggle-edit {
    padding: 0.3em 0.8em;
    background-color: #e0e0e0;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.2s;
}

.toggle-delete:hover {
    background-color: #f44336;
    color: white;
}

.toggle-edit:hover {
    background-color: #4CAF50;
    color: white;
}
</style>
