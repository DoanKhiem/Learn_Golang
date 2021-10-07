<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateNhanvienTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('nhanvien', function (Blueprint $table) {
            $table->id();
            $table->string('name');
            $table->tinyInteger('sex');
            $table->date('birthday');
            $table->string('address');
            $table->integer('cmt');
            $table->string('sdt');
            $table->string('dctamtru');
            $table->string('quoctich');
            $table->text('ghichu');
            $table->unsignedBigInteger('id_phongban');
            $table->foreign('id_phongban')->references('id')->on('phongban');
            $table->unsignedBigInteger('id_chucvu');
            $table->foreign('id_chucvu')->references('id')->on('chucvu');
            $table->unsignedBigInteger('id_hopdonglaodong');
            $table->foreign('id_hopdonglaodong')->references('id')->on('hopdonglaodong');
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('nhanvien');
    }
}
