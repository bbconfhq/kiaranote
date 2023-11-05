USE `kiaranote`;

TRUNCATE TABLE note_hierarchy;
TRUNCATE TABLE shared_note;
TRUNCATE TABLE comment;
TRUNCATE TABLE note;
TRUNCATE TABLE audit_log;
TRUNCATE TABLE user;

-- Insert admin / admin user into the `user` table
INSERT INTO `user` (`username`, `password`, `role`) VALUES
    ('admin', '7666af7301d9f4c41a8cb9b0234ff8f5d544e9dae20d710f77713322e76549e1', 'admin');

-- Insert meaningful sample data into the 'note' table with user_id and is_private set to 0

INSERT INTO `note` (`user_id`, `title`, `content`, `is_private`) VALUES
    (1, '1. Meeting Minutes', 'The meeting minutes for our project discussion held on 2023-11-05. We discussed project goals, milestones, and assigned tasks.', 0),
    (1, '2. Travel Plans', 'Planning a trip to Paris, France. We will visit the Eiffel Tower, Louvre Museum, and enjoy some delicious French cuisine.', 0),
    (1, '3. Recipe Book', 'A collection of my favorite recipes. Today, I am making a classic lasagna with layers of pasta, meat sauce, and cheesy goodness.', 0),
    (1, '4. Bucket List', 'My bucket list of things I want to do in the next 5 years:\n1. Travel to Japan and experience the beauty of cherry blossoms.\n2. Learn a new language, preferably Spanish.', 0),
    (1, '5. Dream Journal', 'Recording my dreams and interpretations.\n\n**Dream 1**:\nI was flying through the sky, soaring above the clouds. It felt incredibly liberating and free.\n\n**Dream 2**:\nI was in a peaceful garden, surrounded by vibrant flowers and calmness.', 0),
    (1, '6. Work Journal', 'Keeping track of my daily work activities and reflections.\n\n**2023-11-01**:\nAttended a team meeting to discuss the upcoming project.\n\n**2023-11-02**:\nPresented my project ideas to the team.', 0),
    (1, '7. Personal Diary', 'A diary of my thoughts and experiences.\n\n**2023-11-01**:\nIt was a beautiful sunny day, and I spent the afternoon at the park, reading my favorite book.\n\n**2023-11-02**:\nHad a long conversation with my friend about life and our dreams.', 0),
    (1, '8. Project Ideas', 'A list of project ideas to explore in the future:\n1. Develop a mobile app for language learning.\n2. Launch an online store for handmade crafts.', 0),
    (1, '9. Nature Photography', 'Capturing the beauty of nature through my lens. Today, I photographed a stunning sunset over the ocean, capturing the vibrant colors and reflections in the water.', 0),
    (1, '10. Fitness Goals', 'Tracking my fitness journey. Completed a 5K run today in under 30 minutes. Feeling strong and motivated to achieve more fitness milestones.', 0),
    (1, '11. Financial Planning', 'Managing my finances and savings. Created a budget for the upcoming month, with a focus on saving for future goals.', 0),
    (1, '12. Gardening Journal', 'Documenting my gardening adventures. Planted a variety of flowers and herbs in my garden today, looking forward to watching them bloom.', 0),
    (1, '13. Book Recommendations', 'Sharing my favorite book recommendations. Just finished reading "The Alchemist" by Paulo Coelho - a truly inspiring and thought-provoking read.', 0),
    (1, '14. Favorite Quotes', 'A collection of quotes that inspire me. "The only way to do great work is to love what you do." - Steve Jobs', 0),
    (1, '15. Home Renovation', 'Planning a home renovation project. Today, I explored new paint colors for the living room and considered furniture upgrades.', 0),
    (1, '16. Learning to Cook', 'Embarking on a culinary journey. Tried making a complex French dish today, and while it was challenging, the results were delicious.', 0),
    (1, '17. Travel Memories', 'Reliving past travel experiences. Reminiscing about the breathtaking views in Santorini and the delicious street food in Bangkok.', 0),
    (1, '18. Artistic Creations', 'Expressing my creativity through art. Today, I painted a vibrant abstract piece with bold colors and dynamic brushstrokes.', 0),
    (1, '19. Tech Enthusiast', 'Exploring the latest in technology. Attended a tech conference and learned about exciting innovations in AI and robotics.', 0),
    (1, '20. Film Reviews', 'Sharing my thoughts on recent films. Just watched "Inception" and was captivated by its mind-bending plot and stunning visuals.', 0),
    (1, '21. Hiking Adventures', 'Documenting my outdoor excursions. Hiked to the summit of a challenging mountain today, and the panoramic views were breathtaking.', 0),
    (1, '22. Music Playlist', 'Curating a playlist of my favorite songs. Listening to a mix of classic rock, indie pop, and jazz to set the mood for the day.', 0),
    (1, '23. Entrepreneurship', 'Pursuing my entrepreneurial dreams. Worked on a business plan for my startup idea and received positive feedback from mentors.', 0),
    (1, '24. Language Learning', 'Embarking on a language learning journey. Practiced conversational Spanish with a language exchange partner today.', 0),
    (1, '25. Meditation Diary', 'Exploring mindfulness and meditation. Todays session brought a deep sense of inner calm and clarity.', 0),
    (1, '26. Culinary Adventures', 'Exploring different cuisines. Tried making sushi at home, and the results were surprisingly delicious.', 0),
    (1, '27. Environmental Awareness', 'Raising awareness about environmental issues. Attended a sustainability workshop and learned about eco-friendly practices.', 0),
    (1, '28. Motivational Quotes', 'A collection of quotes that keep me inspired. "The future belongs to those who believe in the beauty of their dreams." - Eleanor Roosevelt', 0),
    (1, '29. Pet Diary', 'Documenting my pets daily adventures. Today, my dog had a playful day at the park and made new furry friends.', 0),
    (1, '30. Travel Adventure', 'Planning an adventure to explore the Amazon rainforest. Excited about encountering exotic wildlife and experiencing the lush greenery.', 0),
    (1, '31. Movie Night', 'Movie night with friends! Watching a classic comedy and sharing lots of laughter and popcorn.', 0),
    (1, '32. Art Exhibition', 'Visited an art exhibition featuring contemporary art. The abstract paintings and sculptures were thought-provoking.', 0),
    (1, '33. Beach Vacation', 'Looking forward to a beach vacation. Preparing for sun, sea, and relaxation with a good book.', 0),
    (1, '34. Astronomy Hobby', 'Stargazing and observing the night sky. Setting up the telescope to explore distant galaxies and celestial objects.', 0),
    (1, '35. Music Composition', 'Working on a new music composition. Experimenting with melodies and harmonies for a fresh sound.', 0),
    (1, '36. Yoga Retreat', 'Attending a yoga retreat for relaxation and self-discovery. Finding inner peace and balance through daily yoga sessions.', 0),
    (1, '37. Baking Delights', 'Baking a variety of delicious treats. Today, I made a batch of chocolate chip cookies, and the aroma is heavenly.', 0),
    (1, '38. Poetry Writing', 'Expressing emotions through poetry. Writing verses about love, nature, and the beauty of life.', 0),
    (1, '39. Home Movie Night', 'Organizing a cozy home movie night with family. Popcorn, blankets, and a selection of favorite films.', 0),
    (1, '40. Hiking Escapade', 'Planning a hiking escapade to the Grand Canyon. Anticipating breathtaking vistas and a challenging trek.', 0),
    (1, '41. Historical Tour', 'Exploring a historical city on a guided tour. Learning about ancient civilizations and their remarkable achievements.', 0),
    (1, '42. Adventure Photography', 'Documenting my adventure with photography. Capturing the thrill of ziplining through a lush forest.', 0),
    (1, '43. Thrilling Roller Coasters', 'A day at the amusement park riding thrilling roller coasters. Screams of excitement and adrenaline rushes.', 0),
    (1, '44. Gardening Bliss', 'Finding peace in the garden. Planting vibrant flowers and herbs, creating a haven for birds and butterflies.', 0),
    (1, '45. Motivational Journey', 'Embarking on a motivational journey to improve fitness. Tracking progress and celebrating every milestone.', 0),
    (1, '46. Cultural Exploration', 'Immersing in the culture of a foreign country. Savoring local cuisine, music, and traditions.', 0),
    (1, '47. Writing Novels', 'Diving into novel writing. Crafting compelling characters and intriguing plots for a thrilling mystery novel.', 0),
    (1, '48. Interior Design', 'Redesigning my living space. Choosing stylish furniture and decor to create a cozy and inspiring atmosphere.', 0),
    (1, '49. Wildlife Photography', 'Venturing into the wild for wildlife photography. Capturing majestic animals in their natural habitats.', 0),
    (1, '50. Mountain Expedition', 'Preparing for a mountain expedition. Scaling peaks, overcoming challenges, and discovering breathtaking views.', 0),
    (1, '1. Confidential Entry', 'A private note with confidential information that needs to be protected.', 1),
    (1, '2. Sensitive Record', 'This private note contains sensitive data and should not be shared.', 1),
    (1, '3. Personal Diary', 'A personal diary entry meant only for my eyes. No sensitive information included.', 1),
    (1, '4. Project Confidential', 'Confidential project details are documented here. Security is paramount.', 1),
    (1, '5. Financial Privacy', 'Recording financial records and investment strategies. Private and secure.', 1),
    (1, '6. Health Data Secure', 'Keeping sensitive health-related information in this private note for reference.', 1),
    (1, '7. Secure Credentials', 'An encrypted private note with login credentials and passwords. Security is top priority.', 1),
    (1, '8. Personal Reflections', 'Private thoughts and reflections about life and relationships. Not for public sharing.', 1),
    (1, '9. Business Secrets', 'Confidential business meeting minutes and negotiations are documented in this private note.', 1),
    (1, '10. Private Journal', 'A personal diary for private thoughts and experiences that I want to keep secure.', 1);

-- Insert hierarchical data into the `note_hierarchy` table
-- Assume that `note` table has `id` values from 1 to 49
-- Create a tree structure with a minimum depth of 5

-- Depth 1
INSERT INTO `note_hierarchy` (`note_id`, `parent_note_id`, `order`) VALUES
    (1, 1, 1),
    (2, 2, 2);

-- Depth 2
INSERT INTO `note_hierarchy` (`note_id`, `parent_note_id`, `order`) VALUES
    (3, 1, 1),
    (4, 1, 2),
    (5, 2, 1),
    (6, 2, 2);

-- Depth 3
INSERT INTO `note_hierarchy` (`note_id`, `parent_note_id`, `order`) VALUES
    (7, 3, 1),
    (8, 3, 2),
    (9, 4, 1),
    (10, 4, 2),
    (11, 5, 1),
    (12, 5, 2),
    (13, 6, 1),
    (14, 6, 2);

-- Depth 4
INSERT INTO `note_hierarchy` (`note_id`, `parent_note_id`, `order`) VALUES
    (15, 7, 1),
    (16, 7, 2),
    (17, 8, 1),
    (18, 8, 2),
    (19, 9, 1),
    (20, 9, 2),
    (21, 10, 1),
    (22, 10, 2),
    (23, 11, 1),
    (24, 11, 2),
    (25, 12, 1),
    (26, 12, 2),
    (27, 13, 1),
    (28, 13, 2),
    (29, 14, 1),
    (30, 14, 2);

-- Depth 5
INSERT INTO `note_hierarchy` (`note_id`, `parent_note_id`, `order`) VALUES
    (31, 15, 1),
    (32, 15, 2),
    (33, 16, 1),
    (34, 16, 2),
    (35, 17, 1),
    (36, 17, 2),
    (37, 18, 1),
    (38, 18, 2),
    (39, 19, 1),
    (40, 19, 2),
    (41, 20, 1),
    (42, 20, 2),
    (43, 21, 1),
    (44, 21, 2),
    (45, 22, 1),
    (46, 22, 2),
    (47, 23, 1),
    (48, 23, 2),
    (49, 24, 1);

-- Insert hierarchical comment data into the `comment` table, combined and with modified content
-- Assume `user` table has `id` value 1 and `note` table has `id` values from 0 to 49
-- Create a hierarchical comment structure from depth 1 to 5

-- Depth 1
INSERT INTO `comment` (`parent_id`, `user_id`, `note_id`, `content`) VALUES
    (NULL, 1, 1, 'This is a top-level comment about the first note content.'),
    (NULL, 1, 2, 'A comment related to the second note.'),
    (NULL, 1, 2, 'Another top-level comment on a different note.'),
    (NULL, 1, 3, 'A stand-alone comment on the fourth note.');

-- Depth 2
INSERT INTO `comment` (`parent_id`, `user_id`, `note_id`, `content`) VALUES
    (1, 1, 4, 'A reply to the first comment on the first note.'),
    (1, 1, 5, 'An additional reply to the first comment on the first note.'),
    (2, 1, 6, 'A response to the top-level comment on the second note.'),
    (2, 1, 7, 'Yet another response to the top-level comment on the second note.');

-- Depth 3
INSERT INTO `comment` (`parent_id`, `user_id`, `note_id`, `content`) VALUES
    (3, 1, 8, 'A reply to the second comment on the third note.'),
    (3, 1, 9, 'An additional reply to the second comment on the third note.'),
    (4, 1, 10, 'A response to the second comment on the fourth note.'),
    (4, 1, 11, 'Yet another response to the second comment on the fourth note.');

-- Depth 4
INSERT INTO `comment` (`parent_id`, `user_id`, `note_id`, `content`) VALUES
    (5, 1, 12, 'A reply to the third comment on the fifth note.'),
    (5, 1, 13, 'An additional reply to the third comment on the fifth note.'),
    (6, 1, 14, 'A reply to the fourth comment on the sixth note.'),
    (6, 1, 15, 'An additional reply to the fourth comment on the sixth note.');

-- Depth 5
INSERT INTO `comment` (`parent_id`, `user_id`, `note_id`, `content`) VALUES
    (7, 1, 16, 'A reply to the fifth comment on the seventh note.'),
    (7, 1, 17, 'An additional reply to the fifth comment on the seventh note.'),
    (8, 1, 18, 'A response to the third comment on the eighth note.'),
    (8, 1, 19, 'Yet another response to the third comment on the eighth note.');

-- Insert sample data into the `shared_note` table with fixed `id`, `password` set to "password", and `expire_dt` to one week from now
-- Note IDs are updated to 41 through 50

INSERT INTO `shared_note` (`id`, `user_id`, `note_id`, `password`, `expire_dt`) VALUES
    ('ABcd1234efgh', 1, 41, '80fdbe6504c3d2cfb9ee69894cde79ecbf989e1b29c16fdc724cceac1c1e61b0', DATE_ADD(NOW(), INTERVAL 7 DAY)),
    ('Efgh5678ijkl', 1, 42, '80fdbe6504c3d2cfb9ee69894cde79ecbf989e1b29c16fdc724cceac1c1e61b0', DATE_ADD(NOW(), INTERVAL 7 DAY)),
    ('Ijkl9012mnop', 1, 43, '80fdbe6504c3d2cfb9ee69894cde79ecbf989e1b29c16fdc724cceac1c1e61b0', DATE_ADD(NOW(), INTERVAL 7 DAY)),
    ('Mnop3456qrst', 1, 44, '80fdbe6504c3d2cfb9ee69894cde79ecbf989e1b29c16fdc724cceac1c1e61b0', DATE_ADD(NOW(), INTERVAL 7 DAY)),
    ('Qrst7890uvwx', 1, 45, '80fdbe6504c3d2cfb9ee69894cde79ecbf989e1b29c16fdc724cceac1c1e61b0', DATE_ADD(NOW(), INTERVAL 7 DAY)),
    ('Uvwx1234yzab', 1, 46, '80fdbe6504c3d2cfb9ee69894cde79ecbf989e1b29c16fdc724cceac1c1e61b0', DATE_ADD(NOW(), INTERVAL 7 DAY)),
    ('Yzab5678cdef', 1, 47, '80fdbe6504c3d2cfb9ee69894cde79ecbf989e1b29c16fdc724cceac1c1e61b0', DATE_ADD(NOW(), INTERVAL 7 DAY)),
    ('Cdef9012ghij', 1, 48, '80fdbe6504c3d2cfb9ee69894cde79ecbf989e1b29c16fdc724cceac1c1e61b0', DATE_ADD(NOW(), INTERVAL 7 DAY)),
    ('Ghij3456klmn', 1, 49, '80fdbe6504c3d2cfb9ee69894cde79ecbf989e1b29c16fdc724cceac1c1e61b0', DATE_ADD(NOW(), INTERVAL 7 DAY)),
    ('Klmn7890qrst', 1, 50, '80fdbe6504c3d2cfb9ee69894cde79ecbf989e1b29c16fdc724cceac1c1e61b0', DATE_ADD(NOW(), INTERVAL 7 DAY));

