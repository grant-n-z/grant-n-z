import { TestBed } from '@angular/core/testing';

import { UserService } from './user.service';

describe('OperatorLoginService', () => {
  let service: UserService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(UserService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
