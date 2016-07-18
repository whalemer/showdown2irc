// showdown2irc - use Showdown chat with an IRC client
// Copyright (C) 2016 Konrad Borowski
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

type IRCNumeric int

const (
	ErrNoSuchNick         IRCNumeric = 401
	ErrNoSuchServer                  = 402
	ErrNoSuchChannel                 = 403
	ErrCannotSendToChan              = 404
	ErrTooManyChannels               = 405
	ErrWasNoSuchNick                 = 406
	ErrTooManyTargets                = 407
	ErrNoOrigin                      = 409
	ErrNoRecipient                   = 411
	ErrNoTextToSend                  = 412
	ErrNoTopLevel                    = 413
	ErrWildTopLevel                  = 414
	ErrUnknownCommand                = 421
	ErrNoMOTD                        = 422
	ErrNoAdminInfo                   = 423
	ErrFileError                     = 424
	ErrNoNicknameGiven               = 431
	ErrErroneusNickname              = 432
	ErrNicknameInUse                 = 433
	ErrNickCollision                 = 436
	ErrUserNotInChannel              = 441
	ErrNotOnChannel                  = 442
	ErrUserOnChannel                 = 443
	ErrNoLogin                       = 444
	ErrSummonDisabled                = 445
	ErrUsersDisabled                 = 446
	ErrNotRegistered                 = 451
	ErrNeedMoreParams                = 461
	ErrAlreadyRegistered             = 462
	ErrNoPermForHost                 = 463
	ErrPasswdMismatch                = 464
	ErrYouAreBannedCreep             = 465
	ErrKeySet                        = 467
	ErrChannelIsFull                 = 471
	ErrUnknownMode                   = 472
	ErrInviteOnlyChan                = 473
	ErrBannedFromChan                = 474
	ErrBadChannelKey                 = 475
	ErrNoPrivileges                  = 481
	ErrChanOpPrivIsNeeded            = 482
	ErrCannotKillServer              = 483
	ErrNoOperHost                    = 491
	ErrUmodeUnknownFlag              = 501
	ErrUsersDoNotMatch               = 502

	RplWelcome       = 1
	RplYourHost      = 2
	RplCreated       = 3
	RplMyInfo        = 4
	RplBounce        = 5
	RplUserhost      = 302
	RplIson          = 303
	RplAway          = 301
	RplUnaway        = 305
	RplNowAway       = 306
	RplWhoisUser     = 311
	RplWhoisServer   = 312
	RplWhoisOperator = 313
	RplWhoisIdle     = 317
	RplEndOfWhois    = 318
	RplWhoisChannels = 319
	RplWhowasUser    = 314
	RplEndOfWhowas   = 369
	RplListStart     = 321
	RplList          = 322
	RplListEnd       = 323
	RplChannelModes  = 324
	RplNoTopic       = 331
	RplTopic         = 332
	RplInviting      = 341
	RplSummoning     = 342
	RplVersion       = 351
	RplWhoReply      = 352
	RplEndOfWho      = 315
	RplNamesReply    = 353
	RplEndOfNames    = 366
	RplLinks         = 364
	RplEndOfLinks    = 365
	RplBanList       = 367
	RplEndOfBanList  = 368
	RplInfo          = 371
	RplEndOfInfo     = 374
	RplMOTDStart     = 375
	RplMOTD          = 372
	RplEndOfMOTD     = 376
	RplYouAreOper    = 381
	RplRehashing     = 382
	RplTiem          = 391
	RplUsersStart    = 392
	RplUsers         = 393
	RplEndOfUsers    = 394
	RplNoUsers       = 395

	RplTraceLink       = 200
	RplTraceConnecting = 201
	RplTraceHandshake  = 202
	RplTraceUnknown    = 203
	RplTraceOperator   = 204
	RplTraceUser       = 205
	RplTraceServer     = 206
	RplTraceNewType    = 208
	RplTraceLog        = 261

	RplStatsLinkInfo = 211
	RplStatsCommands = 212
	RplEndOfStats    = 219
	RplStatsUptime   = 242
	RplStatsOLine    = 243
	RplUmodeIs       = 221
	RplLuserClient   = 251
	RplLuserOp       = 252
	RplLuserUnknown  = 253
	RplLuserChannels = 254
	RplLuserMe       = 255
	RplAdminMe       = 256
	RplAdminLoc1     = 257
	RplAdminLoc2     = 258
	RplAdminEmail    = 259
)
