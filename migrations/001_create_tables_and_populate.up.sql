CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tweets (
  id INT AUTO_INCREMENT PRIMARY KEY,
  tweet_user INT NOT NULL,
  content VARCHAR(280) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (tweet_user) REFERENCES users(id)
);

CREATE TABLE user_follows (
	userID INT,
  followedUserID INT,
  PRIMARY KEY (userID, followedUserID),
  FOREIGN KEY (userID) REFERENCES users(id),
  FOREIGN KEY (followedUserID) REFERENCES users(id)
);

-- We create these indexes because one requirement was: "La aplicación tiene que estar optimizada para lecturas"
CREATE INDEX idx_tweets_tweet_user ON tweets(tweet_user);
CREATE INDEX idx_user_follows_user_followed ON user_follows(userID, followedUserID);
CREATE INDEX idx_tweets_user_created_at ON tweets(tweet_user, created_at DESC);

INSERT into users (name) values ("Juani");
INSERT into users (name) values ("Feli");
INSERT into users (name) values ("Nico");
INSERT into users (name) values ("Cami");

INSERT into tweets (tweet_user, content) values (1, "Soy Juani y es mi primer tuit :)");
INSERT into tweets (tweet_user, content) values (2, "Primer tuit de feli");
INSERT into tweets (tweet_user, content) values (3, "Nico probando tuiter");
INSERT into tweets (tweet_user, content) values (1, "Soy Juani y es mi segundo tuit :)");
INSERT into tweets (tweet_user, content) values (2, "Segundo tuit de feli");
INSERT into tweets (tweet_user, content) values (3, "Nico sigue probando tuiter");

INSERT into user_follows (userID, followedUserID) values (4, 1);
INSERT into user_follows (userID, followedUserID) values (4, 3);