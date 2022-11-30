import { createRequest } from './request';

export const request = createRequest({ baseURL:"/api" });

export const mockRequest = createRequest({ baseURL: '/mock' });
