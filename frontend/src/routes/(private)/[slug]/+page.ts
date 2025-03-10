import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const prerender = false;

export const load: PageLoad = ({ params }) => {
  if (params.slug) {
    return {
      slug: params.slug,
    };
  }

  error(404, 'Not found');
};