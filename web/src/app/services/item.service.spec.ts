import { TestBed } from '@angular/core/testing';

import { ItemService } from './item.service';
import { HttpClient } from '@angular/common/http';

describe('ItemService', () => {
  let clientSpy: HttpClient;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers:[ItemService, {provide: HttpClient, useValue: clientSpy}]
    })
  });

  it('should be created', () => {
    const service: ItemService = TestBed.get(ItemService);
    expect(service).toBeTruthy();
  });
});
