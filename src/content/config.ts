// 1. Import utilities from `astro:content`
import { z, defineCollection } from 'astro:content';

// 2. Define your collection(s)
const masirCollection = defineCollection({
    type: 'content', // v2.5.0 and later
    schema: z.object({
      title: z.string(),
      author: z.string().optional(),
      description: z.string().optional(),
    }),
});

// 2. Define your collection(s)
const taagcheCollection = defineCollection({
  type: 'content', // v2.5.0 and later
  schema: z.object({
    title: z.string(),
    author: z.string().optional(),
    description: z.string().optional(),
  }),
});

// 3. Export a single `collections` object to register your collection(s)
//    This key should match your collection directory name in "src/content"
export const collections = {
  'masir': masirCollection,
  'taagche': taagcheCollection,
};